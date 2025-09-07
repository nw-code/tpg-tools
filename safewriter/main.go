package main

import (
	"io"
)

type safeWriter struct {
	w io.Writer
	Error error
}

func (s *safeWriter) Write(data []byte) {
	if s.Error != nil {
		return
	}

	_, err := s.w.Write(data)
	if err != nil {
		s.Error = err
	}
}

func write(w io.Writer) error {
	metadata := []byte("hello\n")
	sw := safeWriter{w: w}

	sw.Write(metadata)
	sw.Write(metadata)
	sw.Write(metadata)
	sw.Write(metadata)

	return sw.Error
}
