package storage

import (
	"errors"
	"mime"
	"os"
	"path/filepath"
	"time"
)

// Metadata bundles everything we expose
type Metadata struct {
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	Mime      string    `json:"mime"`
	CreatedAt time.Time `json:"created_at"`
	SHA256    string    `json:"sha256"`
}

// GetMeta builds Metadata for one file.
func GetMeta(dir, filename string) (*Metadata, error) {
	clean := filepath.Base(filepath.Clean(filename))
	if clean == "" {
		return nil, errors.New("invalid filename")
	}
	path := filepath.Join(dir, clean)

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	sha, err := SHA256File(path)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(clean)

	mimeType := mime.TypeByExtension(ext)

	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	return &Metadata{
		Name:      info.Name(),
		Size:      info.Size(),
		Mime:      mimeType,
		SHA256:    sha,
		CreatedAt: info.ModTime(),
	}, nil
}
