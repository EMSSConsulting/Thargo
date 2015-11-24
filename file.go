package thargo

import (
	"archive/tar"
	"io"
	"os"
)

// FileEntry provides a compression entry type which targets a file
// on the local file system.
type FileEntry struct {
  // Path is the relative path under which the file will be stored within the archive.
	Path string
  // Info is the FileInfo object returned from os.Stat for this file.
  // It is used for retrieving the size, permissions and data contained within the file.
	Info os.FileInfo
}

// Header retrieves the archive header for this file entry.
func (f *FileEntry) Header() (*tar.Header, error) {
	header, err := tar.FileInfoHeader(f.Info, f.Info.Name())
	if err != nil {
		return nil, err
	}

	header.Name = f.Path
	return header, nil
}

// Data returns a reader for this file's raw data. You must ensure that
// the returned reader is closed when you are done with it.
func (f *FileEntry) Data() (io.Reader, error) {
	return os.OpenFile(f.Info.Name(), os.O_RDONLY, os.ModePerm)
}
