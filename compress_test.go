package thargo

import (
	"bytes"
	"testing"
)

func TestAddToArchive(t *testing.T) {
	options := *DefaultOptions
	options.CreateIfMissing = true

	buf := new(bytes.Buffer)

	core := NewArchive(buf, &options)

	target := &StringTarget{
		Name:    "test.txt",
		Content: "test",
	}

	err := core.Add(target)
	if err != nil {
		t.Fatal(err)
	}

	reader, err := core.reader()
	if err != nil {
		t.Fatal(err)
	}

	header, _ := reader.Next()
	if header == nil {
		t.Errorf("Expected reader to read a header from the archive")
	}
}

func TestAddIf(t *testing.T) {
	options := *DefaultOptions
	options.CreateIfMissing = true

	buf := new(bytes.Buffer)

	core := NewArchive(buf, &options)

	target := &StringTarget{
		Name:    "test.txt",
		Content: "test",
	}

	err := core.AddIf(target, func(entry Entry) bool {
		return true
	})

	if err != nil {
		t.Fatal(err)
	}

	reader, err := core.reader()
	if err != nil {
		t.Fatal(err)
	}

	header, _ := reader.Next()
	if header == nil {
		t.Errorf("Expected reader to read a header from the archive")
	}

	err = core.AddIf(target, func(entry Entry) bool {
		return false
	})

	if err != nil {
		t.Fatal(err)
	}

	header, _ = reader.Next()
	if header != nil {
		t.Errorf("Expected no additional headers to be available in the archive")
	}
}
