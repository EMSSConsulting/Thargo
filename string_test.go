package thargo

import "testing"

func TestStringTarget(t *testing.T) {
	target := &StringTarget{
		Name:    "test.txt",
		Content: "test",
	}

	entries, err := target.Entries()
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) == 0 {
		t.Fatal("Expected string target to return entries")
	}

	if len(entries) != 1 {
		t.Errorf("Expected string target to return a single entry, got %d instead", len(entries))
	}

	entry := entries[0]

	header, err := entry.Header()
	if err != nil {
		t.Fatal(err)
	}

	if header.Name != "test.txt" {
		t.Errorf("Expected string target header name to be 'test.txt', got '%s' instead", header.Name)
	}

	if header.Mode != 0777 {
		t.Errorf("Expected string target mode to be 0777, got %d instead", header.Mode)
	}

	if header.Size != int64(len("test")) {
		t.Errorf("Expected string target header length to be 4, got %d instead", header.Size)
	}

	data, err := entry.Data()
	if err != nil {
		t.Fatal(err)
	}

	dataOut := make([]byte, header.Size)
	n, err := data.Read(dataOut)
	if err != nil {
		t.Fatal(err)
	}

	dataString := string(dataOut[:n])

	if dataString != "test" {
		t.Errorf("Expected data contents to be 'test', got '%s' instead", dataString)
	}
}
