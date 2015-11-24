package thargo

import (
	"testing"

	"path/filepath"
  "os"
)

func TestNewArchiveFile(t *testing.T) {
  options := *DefaultOptions
  options.CreateIfMissing = true
  
  core, err := NewArchiveFile("test/test.tar.gz", &options)
	if err != nil {
		t.Fatal(err)
	}
  
  defer core.Close()
  defer os.Remove(core.File.Name())

  absPath, err := filepath.Abs("test/test.tar.gz")
	if err != nil {
		t.Fatal(err)
	}

  if core.File.Name() != absPath {
		t.Errorf("Expected NewArchive to expand the archive path to an absolute path.")
	}
}
