package thargo

import (
	"path/filepath"
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

	if header.ChangeTime.IsZero() {
		t.Errorf("Expected entry change time to be non-zero")
	}

	if header.AccessTime.IsZero() {
		t.Errorf("Expected entry access time to be non-zero")
	}
}

func TestFileSystemTargetAbsolutePattern(t *testing.T) {
	absPath, err := filepath.Abs("README.md")
	if err != nil {
		t.Fatal(err)
	}

	target := FileSystemTarget{
		Path:    "./",
		Pattern: absPath,
	}

	entries, err := target.Entries()
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) != 1 {
		t.Error("Expected one entry to be found for README's absolute path")
	}

	header, err := entries[0].Header()
	if err != nil {
		t.Fatal(err)
	}

	if header.Name != "README.md" {
		t.Errorf("Expected entry name to be README.md, got %s instead", header.Name)
	}
}
