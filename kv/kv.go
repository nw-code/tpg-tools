package kv

import (
	"encoding/json"
	"io"
	"os"
)

type store struct {
	data map[string]string
	path string
}

func OpenStore(path string) (*store, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return &store{}, err
	}
	defer f.Close()
	fileinfo, err := f.Stat()
	if err != nil {
		return &store{}, err
	}
	if fileinfo.Size() == 0 {
		f.Write([]byte("{}"))
		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			return &store{}, err
		}
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return &store{}, err
	}
	s := &store{
		data: map[string]string{},
		path: path,
	}
	err = json.Unmarshal(data, &s.data)
	if err != nil {
		s.data = map[string]string{}
		return s, err
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
	data, err := json.Marshal(s.data)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
