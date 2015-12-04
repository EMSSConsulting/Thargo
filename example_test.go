package thargo

import (
	"bytes"
	"fmt"
	"log"
)

// ExampleRoundTrip demonstrates how one would go about first compressing
// a target (in this case, a string) and then extracting that target again
// using the visitor function.
func ExampleRoundTrip() {
	archiveBuffer := new(bytes.Buffer)
	archive := NewArchive(archiveBuffer, nil)

	target := &StringTarget{
		Name:    "test.txt",
		Content: "test string",
	}

	fmt.Printf("Target (%s): %s\n", target.Name, target.Content)

	if _, err := archive.Add(target); err != nil {
		log.Fatal(err)
	}

	if err := archive.Extract(func(entry SaveableEntry) error {
		header, err := entry.Header()
		if err != nil {
			return err
		}

		data, err := entry.Data()
		if err != nil {
			return err
		}

		dataOut := make([]byte, header.Size)
		n, err := data.Read(dataOut)
		if err != nil {
			return err
		}

		extractedString := string(dataOut[:n])

		fmt.Printf("Extract (%s): %s\n", header.Name, extractedString)

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Output:
	// Target (test.txt): test string
	// Extract (test.txt): test string
}
