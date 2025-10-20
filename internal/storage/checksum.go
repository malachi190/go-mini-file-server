package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// SHA256File returns hex-encoded SHA-256 of the file at path.
func SHA256File(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	defer file.Close()

	h := sha256.New()

	if _, err := io.Copy(h, file); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
