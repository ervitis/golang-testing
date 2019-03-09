package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Reader interface {
	ReadData(path string, parser interface{}) error
}

type JsonReader struct{}

func (j *JsonReader) ReadData(path string, parser interface{}) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	if err = json.Unmarshal(b, parser); err != nil {
		return err
	}

	return nil
}