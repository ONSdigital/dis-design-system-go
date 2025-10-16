package helper

import (
	"strconv"

	"github.com/c2h5oh/datasize"
)

func HumanSize(size string) (string, error) {
	if size == "" {
		return "", nil
	}

	s, err := strconv.Atoi(size)
	if err != nil {
		return "", err
	}

	if s < 0 {
		return "", nil
	}

	return datasize.ByteSize(uint64(s)).HumanReadable(), nil
}
