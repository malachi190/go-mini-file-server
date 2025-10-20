package storage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ErrNotFound error is returned when the file is not found
var ErrNotFound = errors.New("file not found")

// Open returns a ReadCloser and the os.FileInfo for a previously stored file
// It sanitizes the name and checks inside dir only
func Open(dir, filename string) (*os.File, os.FileInfo, error) {
	clean := filepath.Base(filepath.Clean(filename))
	path := filepath.Join(dir, clean)

	info, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil, ErrNotFound
		}
		return nil, nil, fmt.Errorf("stat: %w", err)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("open: %w", err)
	}

	return f, info, nil
}
