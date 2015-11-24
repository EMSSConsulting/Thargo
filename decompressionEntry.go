package thargo

import (
	"archive/tar"
	"fmt"
	"io"
)

// DecompressionEntry represents an entry within a tar archive
// which can be decompressed onto the file system.
type DecompressionEntry struct {
	RawEntry
	data io.Closer
}

func newDecompressionEntry(archive *tar.Reader) (*DecompressionEntry, error) {
	header, err := archive.Next()
	if err == io.EOF {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to read header: %s", err)
	}

	if header == nil {
		return nil, nil
	}

	dataReader := &closeableReader{
		reader: io.LimitReader(archive, header.Size),
	}

	return &DecompressionEntry{
		RawEntry: RawEntry{
			RawHeader: header,
			RawData:   dataReader,
		},
		data: dataReader,
	}, nil
}
