package thargo

import (
	"archive/tar"
	"io"
)

// tharWriter is a buffered IO wrapper which provides functionality for writing
// CompressionEntry entities into a tar archive.
type tharWriter struct {
	Writer   *tar.Writer
	Flushers []flushableWriter
	Closers  []closeableWriter
}

type flushableWriter interface {
	io.Writer
	Flush() error
}

type closeableWriter interface {
	io.Writer
	Close() error
}

// Write will write a compression entry into the tar archive
func (w *tharWriter) Write(entry Entry) error {
	header, err := entry.Header()
	if err != nil {
		return err
	}

	err = w.Writer.WriteHeader(header)
	if err != nil {
		return err
	}

	data, err := entry.Data()
	if err != nil {
		return err
	}

	_, err = io.Copy(w.Writer, data)
	if err != nil {
		return err
	}

	return nil
}

func (w *tharWriter) Flush() error {
	for _, writer := range w.Flushers {
		if err := writer.Flush(); err != nil {
			return err
		}
	}

	return nil
}

func (w *tharWriter) Close() error {
	for _, writer := range w.Closers {
		if err := writer.Close(); err != nil {
			return err
		}
	}

	return nil
}
