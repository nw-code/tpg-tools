package kv_test

import (
	"os"
	"testing"

	"github.com/nw-code/tpg-tools/kv"
)

func TestGetReturnsNotOKIfKeyDoesNotExist(t *testing.T) {
	s, err := kv.OpenStore("testdata/non-existent")
	if err != nil {
		t.Fatal(err)
	}
	_, ok := s.Get("foo")
	if ok {
		t.Fatal("unexpected ok")
	}
}

func TestGetReturnsOKWithValueIfKeyExists(t *testing.T) {
	s, err := kv.OpenStore("testdata/non-existent")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "value")
	got, ok := s.Get("key")
	if !ok {
		t.Fatal("not ok")
	}
	if got != "value" {
		t.Errorf("Got %q, want 'value'", got)
	}
}

func TestSetOverwritesExistingKey(t *testing.T) {
	s, err := kv.OpenStore("testdata/non-existent")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "value")
	s.Set("key", "new-value")
	got, ok := s.Get("key")
	if !ok {
		t.Fatal("not ok")
	}
	if got != "new-value" {
		t.Errorf("got %q, want 'new-value'", got)
	}
}

func TestSetToFilePersisted(t *testing.T) {
	path := t.TempDir() + "/kvtest.store"
	s, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	s.Set("A", "1")
	s.Set("B", "2")
	s.Set("C", "3")
	err = s.Save()
	if err != nil {
		t.Fatal(err)
	}
	s2, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	if got, _ := s2.Get("A"); got != "1" {
		t.Errorf("Got A=%s, want A=1", got)
	}
	if got, _ := s2.Get("B"); got != "2" {
		t.Errorf("Got B=%s, want B=2", got)
	}
	if got, _ := s2.Get("C"); got != "3" {
		t.Errorf("Got C=%s, want C=3", got)
	}
}

func TestOpenStore_ErrorsWhenPathUnreadable(t *testing.T) {
	path := t.TempDir() + "/unreadable.store"
	if _, err := os.Create(path); err != nil {
		t.Fatal(err)
	}

	if err := os.Chmod(path, 0o000); err != nil {
		t.Fatal(err)
	}
	_, err := kv.OpenStore(path)
	if err == nil {
		t.Fatal("no error")
	}
}

func TestOpenStore_ReturnsErrorOnInvalidData(t *testing.T) {
	path := t.TempDir() + "/invalid.store"
	if _, err := os.Create(path); err != nil {
		t.Fatal(err)
	}
	_, err := kv.OpenStore(path)
	if err == nil {
		t.Fatal("no error")
	}
}

func TestOpenStore_ErrorsWhenPathUnwritable(t *testing.T) {
	path := "non-existent/unwritable.store"
	s, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Save()
	if err == nil {
		t.Fatal("no error")
	}
}
