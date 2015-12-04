package thargo

// Add will add a compression target to this archive.
func (a *Archive) Add(target Target) (int, error) {
	return a.AddFiltered([]Target{target}, Filter{})
}

// EntryFilterFunc is used in conjunction with AddIf to
// add entries to an archive based on a pre-condition.
type EntryFilterFunc func(entry Entry) bool

// EntriesFilterFunc is used in conjunction with AddAllIf to
// add entries to an archive based on a pre-condition which affects all
// associated entries.
type EntriesFilterFunc func(entries []Entry) bool

// Filter allows you to specify the filter to be used for filtering
// the entries to be added to the archive.
type Filter struct {
	Entry   EntryFilterFunc
	Entries EntriesFilterFunc
}

// AddIf allows you to apply a predicate filter to each
// of the entries which would be included in your archive.
// This can be used to ensure that duplicates are not
// added, if that is a concern, or simply to be notified
// of each file which was added to the archive.
func (a *Archive) AddIf(target Target, predicate EntryFilterFunc) (int, error) {
	return a.AddFiltered([]Target{target}, Filter{
		Entry: predicate,
	})
}

// AddAllIf allows you to conditionally add a group of targets
// and their entries to the archive if they pass the given predicate.
// This can be used to implement things like only updating the archive
// if the source has been updated for example.
func (a *Archive) AddAllIf(targets []Target, predicate EntriesFilterFunc) (int, error) {
	return a.AddFiltered(targets, Filter{
		Entries: predicate,
	})
}

// AddFiltered allows you to use both the EntriesFilterFunc as an "all or nothing"
// predicate as well as the EntryFilterFunc as a per-entry predicate for inclusion
// in the archive.
// This function is intended for use cases where you require the functionality
// available in AddIf and AddAllIf in one combined solution.
func (a *Archive) AddFiltered(targets []Target, filter Filter) (int, error) {
	entries := []Entry{}
	for _, target := range targets {
		e, err := target.Entries()
		if err != nil {
			return 0, err
		}

		entries = append(entries, e...)
	}

	if filter.Entries != nil && !filter.Entries(entries) {
		return 0, nil
	}

	writer, err := a.writer()
	if err != nil {
		return 0, err
	}

	defer writer.Flush()
	defer writer.Close()
  
  added := 0

	for _, entry := range entries {
		if filter.Entry != nil && !filter.Entry(entry) {
			continue
		}

		err := writer.Write(entry)
		if err != nil {
			return 0, err
		}
    
    added++
	}

	return added, nil
}
