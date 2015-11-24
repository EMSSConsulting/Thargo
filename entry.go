package thargo

import (
	"archive/tar"
	"io"
)

// Entry describes the interface which all data sources
// should adhere to in order to support compression within Thargo.
type Entry interface {
	Header() (*tar.Header, error)
	Data() (io.Reader, error)
}
