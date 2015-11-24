package thargo

import (
	"os"
	"path/filepath"
)

// FileSystemTarget provides a compression target for file system entries, which are searched
// for files using a Glob pattern.
type FileSystemTarget struct {
	Path    string
	Pattern string
}

// Entries retrieves the list of compression entries which should be
// included for this target.
func (t *FileSystemTarget) Entries() ([]Entry, error) {
	entries := []Entry{}
	visit := func(path string, f os.FileInfo, err error) error {
		// We don't care about directories, only files
		if f.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(t.Path, path)
		if err != nil {
			return err
		}

		if t.Pattern != "" {
			matched, err := filepath.Match(t.Pattern, relativePath)
			if err != nil {
				return err
			}

			if !matched {
				return nil
			}
		}

		entries = append(entries, &FileEntry{Path: relativePath, Info: f})
		return nil
	}

	err := filepath.Walk(t.Path, visit)

	if err != nil {
		return nil, err
	}

	return entries, nil
}
