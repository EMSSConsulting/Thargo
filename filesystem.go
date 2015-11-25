package thargo

import (
	"os"
	"path"
	"path/filepath"

	"github.com/EMSSConsulting/doublestar"
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
	if filepath.IsAbs(t.Pattern) {
		return t.forAbsolutePath(t.Pattern)
	}
  
	basePath, err := filepath.Abs(t.Path)
	if err != nil {
		return nil, err
	}

	fullPath := path.Join(basePath, filepath.FromSlash(t.Pattern))

	// Check if the file/folder exists, if it does then treat it as
	// an absolute path (possibly enumerating a directory's contents)
	_, err = os.Stat(fullPath)
	if err != nil {
		return t.forGlobPath(basePath, t.Pattern)
	}

	return t.forAbsolutePath(fullPath)
}

func (t *FileSystemTarget) forAbsolutePath(path string) ([]Entry, error) {
	entries := []Entry{}

	basePath, err := filepath.Abs(t.Path)
	if err != nil {
		return nil, err
	}

	relativePath, err := filepath.Rel(basePath, path)
	if err != nil {
		relativePath = path
	}

	f, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if err != nil {
		return entries, nil
	}

	// If we are given a directory, enumerate its contents recursively
	if f.IsDir() {
		return t.forGlobPath(path, "**")
	}

	entries = append(entries, &FileEntry{Name: relativePath, Path: path, Info: f})
	return entries, nil
}

func (t *FileSystemTarget) forGlobPath(basePath, glob string) ([]Entry, error) {
	entries := []Entry{}

	matches, err := doublestar.Glob(basePath, glob)
	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		relativePath, err := filepath.Rel(basePath, match)
		if err != nil {
			return nil, err
		}

		f, err := os.Stat(match)
		if err != nil {
			return nil, err
		}

		entries = append(entries, &FileEntry{Name: relativePath, Path: match, Info: f})
	}

	return entries, nil
}
