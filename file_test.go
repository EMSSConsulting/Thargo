package thargo

import (
	"os"
	"path/filepath"
	"testing"
)

func testGetFile(path string) (*FileEntry, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}

	return &FileEntry{
		Path: path,
		Info: info,
	}, nil
}

func TestFileHeader(t *testing.T) {
	file, err := testGetFile("README.md")
	if err != nil {
		t.Fatal(err)
	}

	header, err := file.Header()
	if err != nil {
		t.Fatal(err)
	}

	if header.Name != "README.md" {
		t.Errorf("Expected name to be the relative path.")
	}

	if header.Size <= 0 {
		t.Errorf("Expected file size to be set correctly")
	}

	if header.Mode == 0 {
		t.Errorf("Expected file mode to be set")
	}
}
