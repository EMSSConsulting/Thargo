package thargo

import (
	"testing"

	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
)

func TestWriter(t *testing.T) {
	buf := new(bytes.Buffer)

	options := *DefaultOptions
	core := NewArchive(buf, &options)

	writer, err := core.writer()
	if err != nil {
		t.Fatal(err)
	}

	dataBuffer := new(bytes.Buffer)
	dataBuffer.WriteString("test")

	err = writer.Write(&RawEntry{
		RawHeader: &tar.Header{
			Name: "test.txt",
			Size: 4,
			Mode: 0777,
		},
		RawData: dataBuffer,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Flush()
	if err != nil {
		t.Error(err)
	}

	if buf.Len() == 0 {
		t.Error("Expected buffer to have been written to")
	}
}

func TestReader(t *testing.T) {
	buf := new(bytes.Buffer)

	options := *DefaultOptions
	core := NewArchive(buf, &options)

	gw := gzip.NewWriter(buf)
	tw := tar.NewWriter(gw)
	if err := tw.WriteHeader(&tar.Header{
		Name: "test.txt",
		Size: 4,
		Mode: 0777,
	}); err != nil {
		t.Fatal(err)
	}

	if _, err := tw.Write([]byte("test")); err != nil {
		t.Fatal(err)
	}

	tw.Flush()
	gw.Flush()

	r, err := core.reader()
	if err != nil {
		t.Fatal(err)
	}

	header, err := r.Next()
	if err != nil {
		t.Fatal(err)
	}

	if header.Name != "test.txt" {
		t.Error("Expected header name to be test.txt")
	}

	if header.Size != 4 {
		t.Error("Expected header size to be 4")
	}

	if header.Mode != 0777 {
		t.Error("Expected header mode to be 0777")
	}

	header, err = r.Next()
	if header != nil {
		t.Error("Expected header to be nil at the end of the archive")
	}

	if header != nil && err != io.EOF {
		t.Error("Expected EOF error to be returned at the end of the archive")
	}
}
