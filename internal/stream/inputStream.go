package stream

import (
	"io"
	"os"
)

type Input struct {
	Reader io.Reader
	File   *os.File
}

func NewInputStream(filename string) (*Input, error) {

	if filename == "" {
		return &Input{Reader: os.Stdin}, nil
	}

	infile, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	return &Input{Reader: infile, File: infile}, nil
}

func (s *Input) Close() {
	if s.File != nil {
		s.File.Close()
	}
}
