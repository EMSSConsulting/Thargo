package thargo

// DecompressionVisitor is the type of the function called for each entry within
// the archive. It will be called as each entry is enumerated and the data reader
// is guaranteed to be available to read data correctly under the provision that
// Read is called synchronously within the visitor function, or one of its synchronous
// calls.
// Once you have returned from this function, the data reader will be closed and you
// will no longer be able to read data from it.
type DecompressionVisitor func(entry SaveableEntry) error

// Extract will decompress this archive into the specified destination
// directory.
func (a *Archive) Extract(visit DecompressionVisitor) error {
	reader, err := a.reader()
	if err != nil {
		return err
	}

	for {
		entry, err := newDecompressionEntry(reader)
		if err != nil {
			return err
		}

		if entry == nil {
			return nil
		}

		err = visit(entry)
		entry.data.Close()

		if err != nil {
			return err
		}

	}
}
