package thargo

import (
	"bytes"
	"testing"
)

func TestDecompression(t *testing.T) {
	options := *DefaultOptions
	options.CreateIfMissing = true

	buf := new(bytes.Buffer)

	core := NewArchive(buf, &options)

	target := &StringTarget{
		Name:    "test.txt",
		Content: "test",
	}

	added, err := core.Add(target)
	if err != nil {
		t.Fatal(err)
	}
  
  if added != 1 {
    t.Errorf("Expected one entry to be added to the archive")
  }

	wasExtracted := true

	err = core.Extract(func(entry SaveableEntry) error {
		wasExtracted = true

		header, err := entry.Header()
		if err != nil {
			t.Fatal(err)
		}

		if header.Name != "test.txt" {
			t.Errorf("Expected extracted item header to have name: test.txt, got %s", header.Name)
		}

		if header.Size != 4 {
			t.Errorf("Expected extracted item to have length of 4, got %d", header.Size)
		}

		data, err := entry.Data()
		if err != nil {
			t.Fatal(err)
		}

		dataOut := make([]byte, 4)
		written, err := data.Read(dataOut)
		if written != 4 {
			t.Errorf("Expected data to have length of 4, got %d", written)
		}

		if err != nil {
			t.Fatal(err)
		}

		if string(dataOut) != "test" {
			t.Errorf("Expected data to be 'test', got '%s' instead", string(dataOut))
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	if !wasExtracted {
		t.Error("Expected file to be extracted, it was not")
	}
}
