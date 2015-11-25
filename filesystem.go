package thargo

import (
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar"
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

	basePath, err := filepath.Abs(t.Path)
	if err != nil {
		return nil, err
	}

	if filepath.IsAbs(t.Pattern) {
		relativePath, err := filepath.Rel(basePath, t.Pattern)
		if err != nil {
			relativePath = t.Pattern
		}

		f, err := os.Stat(t.Pattern)
		if err != nil && !os.IsNotExist(err) {
			return nil, err
		}

		if err == nil {
			entries = append(entries, &FileEntry{Name: relativePath, Path: t.Pattern, Info: f})
		}

		return entries, nil
	}

	visit := func(path string, f os.FileInfo, err error) error {
		// We don't care about directories, only files
		if f.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(basePath, path)
		if err != nil {
			return err
		}

		if t.Pattern != "" {
			if t.Pattern != relativePath {
				matched, err := doublestar.Match(t.Pattern, filepath.ToSlash(relativePath))
				if err != nil {
					return err
				}

				if !matched {
					return nil
				}
			}
		}

		entries = append(entries, &FileEntry{Name: relativePath, Path: path, Info: f})
		return nil
	}

	if err := filepath.Walk(basePath, visit); err != nil {
		return nil, err
	}

	return entries, nil
}
