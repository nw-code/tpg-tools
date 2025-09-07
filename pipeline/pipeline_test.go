package pipeline_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nw-code/tpg-tools/pipeline"
)

func TestStdoutPrintsMessageToOutput(t *testing.T) {
	want := "hello, world!"
	p := pipeline.FromString(want)
	buf := new(bytes.Buffer)
	p.Output = buf
	p.Stdout()
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got := buf.String()
	if !cmp.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestStdoutPrintsNothingOnError(t *testing.T) {
	p := pipeline.FromString("hello, world!")
	buf := new(bytes.Buffer)
	p.Output = buf
	p.Error = errors.New("some error")
	p.Stdout()
	got := buf.String()
	if got != "" {
		t.Errorf("want no output from Stdout after error, but got %q", got)
	}
}

func TestFromFile_ReadsAllDataFromFile(t *testing.T) {
	want := []byte("hello, world\n")
	p := pipeline.FromFile("testdata/hello.txt")
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got, err := io.ReadAll(p.Reader)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFromFile_SetsErrorGivenNonexistentFile(t *testing.T) {
	p := pipeline.FromFile("testdata/missing.txt")
	if p.Error == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestStringReturnsPipeContents(t *testing.T) {
	want := "hello, world\n"
	p := pipeline.FromString(want)
	got, err := p.String()
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestStringReturnsErrorWhenPipeErrorSet(t *testing.T) {
	p := pipeline.FromFile("testdata/missing.txt")
	_, err := p.String()
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestColumn_ExtractsColumn2(t *testing.T) {
	input := "1 2 3\n1 2 3\n1 2 3\n"
	want := "2\n2\n2\n"
	p := pipeline.FromString(input)
	//p.Column(2)
	//data, err := io.ReadAll(p.Reader)
	got, err := p.Column(2).String()
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(got, want) {
		t.Errorf(cmp.Diff(got, want))
	}
}

func TestColumnProducesNothingWhenPipeErrorSet(t *testing.T) {
	p := pipeline.FromString("1 2 3\n")
	p.Error = errors.New("foobar")
	data, err := io.ReadAll(p.Column(1).Reader)
	if err != nil {
		t.Error(err)
	}
	if len(data) > 0 {
		t.Errorf("want no output from Column after error, but got %q", data)
	}
}

func TestColumnSetsErrorAndProducesNothingGivenInvalidArg(t *testing.T) {
	p := pipeline.FromString("1 2 3\n1 2 3\n1 2 3\n")
	p.Column(-1)
	if p.Error == nil {
		t.Error("want error on non-positive Column, but got nil")
	}
	data, err := io.ReadAll(p.Column(1).Reader)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 {
		t.Errorf("want no output from Column after error, but got %q", data)
	}
}
