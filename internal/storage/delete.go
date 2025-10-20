package storage

import (
	"errors"
	"os"
	"path/filepath"
)

// Delete removes a file from the directory; It returns ErrNotFound if the file is not found
func Delete(dir, filename string) error {
	clean := filepath.Base(filepath.Clean(filename))

	if clean == "" || clean == "." {
		return errors.New("invalid file name")
	}

	path := filepath.Join(dir, clean)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrNotFound
	}

	return os.Remove(path)
}
