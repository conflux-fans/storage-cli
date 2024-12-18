package zkclient

import (
	"fmt"
	"time"
)

func u64ToBytes(n uint64) [8]byte {
	var result [8]byte
	for i := 0; i < 8; i++ {
		result[i] = byte((n >> (i * 8)) & 0xFF)
	}
	return result
}

func fillBytes(source []byte, target []byte) error {
	if len(source) > len(target) {
		return fmt.Errorf("source is longer than target")
	}

	copy(target, source)
	return nil
}

func PaddingLeft0(source []byte, size int) []byte {
	pad := make([]byte, size)
	copy(pad[len(pad)-len(source):], source)
	return pad
}

// return timestamp, input format is yearmonthdate, such as 20240101
func MustParseBirthDate(val string) uint64 {
	t, err := time.Parse("20060102", val)
	if err != nil {
		panic(err)
	}
	return uint64(t.Unix())
}

func BoolToUnit(b bool) uint {
	if b {
		return 1
	}
	return 0
}

func PadToSector(input []byte) []byte {
	sector := make([]byte, 256)
	copy(sector, input)
	return sector
}

func BoolsToUint64(values []bool) uint64 {
	var result uint64
	for _, v := range values {
		// result += uint64(int(BoolToUnit(v)) >> i)
		result = result<<1 | uint64(BoolToUnit(v))
	}
	return result
}

func InvertBools(values []bool) []bool {
	result := make([]bool, len(values))
	for i, v := range values {
		result[i] = !v
	}
	return result
}

func ReverseBools(values []bool) []bool {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return values
}

func stringToByte16(s []byte) [16]byte {
	var byteArray [16]byte
	copy(byteArray[:], s)
	return byteArray
}
