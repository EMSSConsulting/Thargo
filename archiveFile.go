package thargo

import (
	"os"
	"path/filepath"
)

// ArchiveFile represents a tar archive which stores its data on the local filesystem.
// It provides the Close method which allows you to close the file handle when you are
// finished working with the file.
type ArchiveFile struct {
	Archive
	File *os.File
}

// NewArchiveFile creates a new tar archive at the specified location
// on disk.
func NewArchiveFile(path string, options *Options) (*ArchiveFile, error) {
	if options == nil {
		options = DefaultOptions
	}

	absPath, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
  
  // Create the directory necessary to hold the archive file
  absDirPath := filepath.Dir(absPath)
  err = os.MkdirAll(absDirPath, os.ModePerm | os.ModeDir)
  if err != nil {
    return nil, err
  }

  // Create the archive file, or open it if it exists already (depening on options)
	flags := os.O_RDWR
	if options.CreateIfMissing {
		flags = flags | os.O_CREATE
	}
	file, err := os.OpenFile(absPath, flags, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &ArchiveFile{
		Archive: Archive{
			Stream:  file,
			Options: *options,
		},
		File: file,
	}, nil
}

// Close will close the handle to this archive file. It is imperative that you
// call this once you are finished working with an archive file, failure to do so
// may prevent the file from being opened until your program exits.
// Once you have closed the file, you will be unable to read or write to it again.
func (a *ArchiveFile) Close() error {
	return a.File.Close()
}
