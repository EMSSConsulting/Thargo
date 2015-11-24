// Package thargo provides an easy to use API for archiving and compressing
// various data structures using a combination of the tar and gzip formats.
// Thargo provides a higher level abstraction than that offered by the built
// in archive/tar and compress/gzip packages, treating compression as the
// inclusion of a number of entries (provided by various compression targets)
// into an archive. This allows for easy expansion to support various input
// and output targets, while still making their use very straightforward,
// with a minimum of boilerplate being used.
package thargo

import (
	"archive/tar"
	"compress/gzip"
	"io"
)

// Archive is a tar archiver library for Go which abstracts the process
// of archiving and extracting file system structures.
type Archive struct {
	Stream  io.ReadWriter
	Options Options
}

// NewArchive creates a new Thargo archive for you using the provided
// stream.
func NewArchive(stream io.ReadWriter, options *Options) *Archive {
	if options == nil {
		options = DefaultOptions
	}

	return &Archive{
		Stream:  stream,
		Options: *options,
	}
}

func (a *Archive) reader() (*tar.Reader, error) {
	reader := io.Reader(a.Stream)

	if a.Options.GZip {
		gr, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}

		reader = gr
	}

	return tar.NewReader(reader), nil
}

func (a *Archive) writer() (*tharWriter, error) {
	writer := io.Writer(a.Stream)

	flushers := []flushableWriter{}
	closers := []closeableWriter{}

	if a.Options.GZip {
		if a.Options.GZipLevel > 0 {
			gw, err := gzip.NewWriterLevel(writer, a.Options.GZipLevel)
			if err != nil {
				return nil, err
			}

			flushers = append([]flushableWriter{gw}, flushers...)
			closers = append([]closeableWriter{gw}, closers...)
			writer = gw
		} else {
			writer = gzip.NewWriter(writer)
		}
	}

	tw := tar.NewWriter(writer)
	flushers = append([]flushableWriter{tw}, flushers...)

	return &tharWriter{
		Writer:   tw,
		Flushers: flushers,
		Closers:  closers,
	}, nil
}
