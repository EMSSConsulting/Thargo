package thargo

// Add will add a compression target to this archive.
func (a *Archive) Add(target Target) error {
	writer, err := a.writer()
	if err != nil {
		return err
	}

	defer writer.Flush()
	defer writer.Close()

	entries, err := target.Entries()
	if err != nil {
		return err
	}

	for _, entry := range entries {
		err = writer.Write(entry)
		if err != nil {
			return err
		}
	}

	return nil
}
