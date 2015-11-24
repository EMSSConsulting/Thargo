package thargo

// Add will add a compression target to this archive.
func (a *Archive) Add(target Target) error {
	return a.AddIf(target, func(entry Entry) bool {
    return true
  })
}

// EntryFilterFunc is used in conjunction with AddIf to
// add entries to an archive based on a pre-condition.
type EntryFilterFunc func(entry Entry) bool

// AddIf allows you to apply a predicate filter to each
// of the entries which would be included in your archive.
// This can be used to ensure that duplicates are not
// added, if that is a concern, or simply to be notified
// of each file which was added to the archive.
func (a *Archive) AddIf(target Target, predicate EntryFilterFunc) error {
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
    if predicate(entry) {
      err = writer.Write(entry)
      if err != nil {
        return err
      }
    }
	}

	return nil
}