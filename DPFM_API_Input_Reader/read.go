package dpfm_api_input_reader

import (
	"encoding/json"
	"os"

	"golang.org/x/xerrors"
)

type FileReader struct{}

func NewFileReader() *FileReader {
	return &FileReader{}
}

func (FileReader) ReadInput(data []byte) (*Request, error) {
	t := Request{}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, xerrors.Errorf("input file format is bad: %w", err)
	}
	return &t, nil
}

func (FileReader) GetFilePointer(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDONLY, 0)
}
