package thargo

import (
	"fmt"
	"io"
)

type closeableReader struct {
	reader io.Reader
}

func (r *closeableReader) Read(p []byte) (int, error) {
	if r.reader == nil {
		return 0, fmt.Errorf("reader has been closed")
	}

	return r.reader.Read(p)
}

func (r *closeableReader) Close() error {
	r.reader = nil
	return nil
}
