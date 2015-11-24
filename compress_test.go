package thargo

import (
	"testing"
  "bytes"
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
}
