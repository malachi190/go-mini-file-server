package storage

import (
	"os"
	"path/filepath"
	"sort"
	"time"
)

// FileInfo is what we expose to the server
type FileInfo struct {
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	SHA256    string    `json:"sha256"`
	CreatedAt time.Time `json:"created_at"`
}

// List reads the upload directory and returns a sorted slice
func List(dir string) ([]FileInfo, error) {
	entries, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	out := make([]FileInfo, 0, len(entries))

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		info, err := e.Info()

		if err != nil {
			continue // skip problematic entry
		}

		path := filepath.Join(dir, info.Name())

		sha256, err := SHA256File(path)

		if err != nil {
			continue
		}

		out = append(out, FileInfo{
			Name:      info.Name(),
			Size:      info.Size(),
			SHA256:    sha256,
			CreatedAt: info.ModTime(),
		})
	}

	// sort in deterministic order
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}
