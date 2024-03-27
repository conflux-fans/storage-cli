package core

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/zero-gravity-labs/zerog-storage-client/core"
)

const (
	CHUNK_SIZE     = 4096
	VALUE_MAX_SIZE = CHUNK_SIZE * 100
)

// append source file to dest name
func AppendData(name string, data string, force bool) error {
	if len(data) > VALUE_MAX_SIZE {
		return errors.New("Exceed max size once uploadable")
	}
	logrus.WithField("name", name).Info("Start append content")

	// split content to chunks
	var chunks []string
	for i := 0; i < len(data); i += CHUNK_SIZE {
		end := lo.Min([]int{(i + 1) * CHUNK_SIZE, len(data)})
		chunks = append(chunks, data[i*CHUNK_SIZE:end])
	}

	meta, err := GetContentMetadata(name)
	if err != nil {
		if !(err == ERR_UNEXIST_CONTENT && force) {
			return err
		}
	}
	if meta == nil {
		meta = &ContentMetadata{
			LineSizeKey: keyLineCount(name),
			LineKeys:    []string{},
			LineSize:    0,
		}
	}

	batcher := defaultKvClientForPut.Batcher()
	batcher.Set(STREAM_FILE,
		[]byte(meta.LineSizeKey),
		[]byte(fmt.Sprintf("%d", meta.LineSize+len(chunks))),
	)
	for i, chunk := range chunks {
		batcher.Set(STREAM_FILE,
			[]byte(keyLineIndex(name, meta.LineSize+i)),
			[]byte(chunk),
		)
	}

	err = batcher.Exec()
	if err != nil {
		return errors.WithMessage(err, "Failed to execute batcher")
	}

	logrus.WithField("name", name).WithField("line", len(chunks)).Info("Append content completed")

	return nil
}

func AppendFromFile(name string, filePath string, force bool) error {
	f, err := openFile(filePath)
	if err != nil {
		return err
	}

	// split by VALUE_MAX_SIZE
	for {
		buffer := make([]byte, 1024)
		n, err := f.Read(buffer)
		if err != nil {
			return err
		}
		if n == 0 {
			return nil
		}

		if err = AppendData(name, string(buffer[:n]), force); err != nil {
			return err
		}
	}
}

func openFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return nil, core.ErrFileRequired
	}

	if info.Size() == 0 {
		return nil, core.ErrFileEmpty
	}

	if info.Size() > VALUE_MAX_SIZE {
		return nil, fmt.Errorf("file size exceeds maximum size %d", VALUE_MAX_SIZE)
	}
	return file, nil
}

func keyLineCount(name string) string {
	return fmt.Sprintf("%s:line", name)
}

func keyLineIndex(name string, index int) string {
	return fmt.Sprintf("%s:%d", name, index)
}

// // Note: useless
// func AppendFileKeyToDb(filepath string, name string) error {
// 	// save db
// 	fileNameKey := db.KeyFileName(name)
// 	value, err := db.GetDB().Get([]byte(fileNameKey), nil)
// 	if err != nil {
// 		return errors.WithMessagef(err, "Failed to query %s", name)
// 	}

// 	var roots []common.Hash
// 	if err := json.Unmarshal(value, &roots); err != nil {
// 		return err
// 	}

// 	rootHash, err := GetRootHash(filepath)
// 	if err != nil {
// 		return err
// 	}

// 	j, err := json.Marshal(append(roots, rootHash))
// 	if err != nil {
// 		return err
// 	}

// 	err = db.GetDB().Put([]byte(name), j, nil)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }