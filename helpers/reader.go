package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Reader interface {
	ReadData(path string) ([]byte, error)
}

type JsonReader struct{}

func (j *JsonReader) ReadData(path string) ([]byte, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	b, _ := ioutil.ReadAll(f)

	return b, nil
}