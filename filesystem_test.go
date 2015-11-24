package thargo

import (
	"testing"
)

func TestFileSystemTarget(t *testing.T) {
	target := FileSystemTarget{
		Path:    "./",
		Pattern: "README*",
	}

	entries, err := target.Entries()
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) != 1 {
		t.Error("Expected one entry to be found for ./, README*")
	}

	header, err := entries[0].Header()
	if err != nil {
		t.Fatal(err)
	}

	if header.Name != "README.md" {
		t.Errorf("Expected entry name to be README.md, got %s instead", header.Name)
	}
}
