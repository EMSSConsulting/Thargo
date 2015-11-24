package thargo

import (
	"archive/tar"
	"io"
	"os"
)

// FileEntry provides a compression entry type which targets a file
// on the local file system.
type FileEntry struct {
	path string
	info os.FileInfo
}

// Header retrieves the archive header for this file entry.
func (f *FileEntry) Header() (*tar.Header, error) {
	header, err := tar.FileInfoHeader(f.info, f.info.Name())
	if err != nil {
		return nil, err
	}

	header.Name = f.path
	return header, nil
}

// Data returns a reader for this file's raw data. You must ensure that
// the returned reader is closed when you are done with it.
func (f *FileEntry) Data() (io.Reader, error) {
	return os.OpenFile(f.info.Name(), os.O_RDONLY, os.ModePerm)
}
