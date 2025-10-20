package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CreateDir creates the directory if missing
func CreateDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

// Save writes the uploaded file to a disk with the original name
func Save(dir string, filename string, src io.Reader) (int64, error) {
	path := filepath.Join(dir, filepath.Base(filename))

	dest, err := os.Create(path)

	if err != nil {
		return 0, fmt.Errorf("create file: %v", err)
	}

	defer dest.Close()

	written, err := io.Copy(dest, src)

	if err != nil {
		return written, fmt.Errorf("write file: %w", err)
	}

	return written, nil
}
