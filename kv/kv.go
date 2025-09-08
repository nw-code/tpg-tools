package kv

import (
	"encoding/gob"
	"errors"
	"io/fs"
	"os"
)

type store struct {
	data map[string]string
	path string
}

func OpenStore(path string) (*store, error) {
	s := &store{
		data: map[string]string{},
		path: path,
	}
	f, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = gob.NewDecoder(f).Decode(&s.data)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *store) Get(key string) (string, bool) {
	v, ok := s.data[key]
	return v, ok
}

func (s *store) Set(key, value string) {
	s.data[key] = value
}

func (s *store) Save() error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()
	return gob.NewEncoder(f).Encode(s.data)
}
