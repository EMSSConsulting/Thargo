package thargo

import (
  "testing"
  "bytes"
  "os"
)

func TestDecompressionEntry(t *testing.T) {
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
		t.Fatalf("Failed to compress file: %s", err)
	}
  
  wasExtracted := false
  
  err = core.Extract(func(entry SaveableEntry) error {
    wasExtracted = true
    
    if err := entry.Save("test/"); err != nil {
      t.Fatalf("Failed to save file: %s", err)
    }
    
    defer os.Remove("test/test.txt")
    
    fileInfo, err := os.Stat("test/test.txt")
    if err != nil {
      t.Fatalf("Failed to stat file: %s", err)
    }
    
    if fileInfo.Size() != 4 {
      t.Errorf("Expected saved file size to be 4, got %d instead", fileInfo.Size())
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