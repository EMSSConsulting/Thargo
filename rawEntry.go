package thargo

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

// RawEntry can be used in situations where you do not wish
// to make use of a higher-level abstraction, or wish to formulate the
// header yourself.
type RawEntry struct {
	RawHeader *tar.Header
	RawData   io.Reader
}

// Header retrieves the compression entry's tar archive header.
func (e *RawEntry) Header() (*tar.Header, error) {
	return e.RawHeader, nil
}

// Data retreives an io.Reader which supplies the data to be written for this entry.
func (e *RawEntry) Data() (io.Reader, error) {
	return e.RawData, nil
}

// Save will save this decompression entry to file within the destination
// directory provided, making use of the correct relative paths etc.
func (e *RawEntry) Save(destination string) error {
	destination, err := filepath.Abs(destination)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(destination, e.RawHeader.Name)
	fullDirPath := filepath.Dir(fullPath)

	err = os.MkdirAll(fullDirPath, os.ModePerm|os.ModeDir)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(e.RawHeader.Mode))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, e.RawData)
	if err != nil {
		return err
	}

	return nil
}
