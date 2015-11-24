package thargo

// CompressionTarget represents the interface which all high-level
// targets should implement in order to support being passed to Thargo.
type Target interface {
	Entries() ([]Entry, error)
}
