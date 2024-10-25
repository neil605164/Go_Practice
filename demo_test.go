package main

import (
	"testing"

	"github.com/gin-gonic/gin/binding"
)

type ByteID []byte

// implements encoding.TextUnmarshaler
func (id *ByteID) UnmarshalText(data []byte) error {
	*id = []byte("called")
	return nil
}

type bytePath struct {
	ID ByteID `uri:"id"`
}

func TestURIBindingByteID(t *testing.T) {
	m := map[string][]string{
		"id": {"asdf"},
	}

	var path bytePath
	err := binding.Uri.BindUri(m, &path)
	if err != nil {
		t.Fatal(err)
	}
	if string(path.ID) != "called" {
		t.Fatal("UnmarshalText not called")
	}
}

type TextID string

// implements encoding.TextUnmarshaler
func (id *TextID) UnmarshalText(data []byte) error {
	*id = "called"
	return nil
}

type textPath struct {
	ID TextID `uri:"id"`
}

func TestURIBindingTextID(t *testing.T) {
	m := map[string][]string{
		"id": {"asdf"},
	}

	var path textPath
	err := binding.Uri.BindUri(m, &path)
	if err != nil {
		t.Fatal(err)
	}
	if string(path.ID) != "called" {
		t.Fatal("UnmarshalText not called")
	}
}
