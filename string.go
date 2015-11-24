package thargo

import (
	"archive/tar"
	"strings"
)

// StringTarget provides a data source target for a string to be included in
// your archive under the given name.
type StringTarget struct {
	Name    string
	Content string
}

// Entries retrieves the list of compression entries which should be
// included for this target.
func (t *StringTarget) Entries() ([]Entry, error) {
	data := strings.NewReader(t.Content)
	return []Entry{
		&RawEntry{
			RawHeader: &tar.Header{
				Name: t.Name,
				Mode: 0777,
				Size: int64(data.Len()),
			},
			RawData: data,
		},
	}, nil
}
