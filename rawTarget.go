package thargo

// RawTarget provides a low level target implementation allowing
// you to manually specify the list of entries you wish to include
// within it.
type RawTarget struct {
	RawEntries []Entry
}

// Entries retrieves the list of compression entries which should be
// included for this target.
func (t *RawTarget) Entries() ([]Entry, error) {
	return t.RawEntries, nil
}
