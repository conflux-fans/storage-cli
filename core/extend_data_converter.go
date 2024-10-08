package core

import (
	ccore "github.com/0glabs/0g-storage-client/core"
	"github.com/conflux-fans/storage-cli/config"
	"github.com/conflux-fans/storage-cli/constants/enums"
)

type ExtendDataConverter struct {
}

var extendDataConverter ExtendDataConverter

func DefaultExtendDataConverter() *ExtendDataConverter {
	return &extendDataConverter
}

func (p *ExtendDataConverter) ByContent(data []byte) (enums.ExtendDataType, ccore.IterableData, error) {
	_data, err := ccore.NewDataInMemory(data)
	if err != nil {
		return enums.ExtendDataType(-1), nil, err
	}
	return getExtendType(_data.Size()), _data, nil
}

func (p *ExtendDataConverter) ByFile(filePath string) (enums.ExtendDataType, ccore.IterableData, error) {
	f, err := ccore.Open(filePath)
	if err != nil {
		return enums.ExtendDataType(-1), nil, err
	}
	return getExtendType(f.Size()), f, nil
}

func getExtendType(size int64) enums.ExtendDataType {
	if size <= config.Get().ExtendData.TextMaxSize {
		return enums.EXTEND_DATA_TEXT
	}
	return enums.EXTEND_DATA_POINTER
}
