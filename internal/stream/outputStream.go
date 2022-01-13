package stream

import (
	"io"
	"os"
)

type Output struct {
	Writer io.Writer
	File   *os.File
}

func NewOutputStream(filename string) (*Output, error) {

	if filename == "" {
		return &Output{Writer: os.Stdout}, nil
	}

	outfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return nil, err
	}

	return &Output{Writer: outfile, File: outfile}, nil
}

func (s *Output) Close() {
	if s.File != nil {
		s.File.Close()
	}
}
