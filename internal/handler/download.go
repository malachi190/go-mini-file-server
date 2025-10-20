package handler

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/malachi190/go-file-server/internal/storage"
)

// Download returns an http.HandlerFunc for GET /files/{filename}
func Download(uploadDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract filename from path
		filename := filepath.Base(r.URL.Path) // removes leading slash

		if filename == "" || filename == "files" {
			http.Error(w, "missing filename", http.StatusBadRequest)
			return
		}

		file, info, err := storage.Open(uploadDir, filename)

		if err != nil {
			if err == storage.ErrNotFound {
				http.Error(w, "file not found", http.StatusNotFound)
				return
			}
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		defer file.Close()

		// Go detects MIME; fallback to octet stream
		w.Header().Set("Content-Type", detectMIME(info.Name()))
		w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))
		// Suggest download with original name
		w.Header().Set("Content-Disposition", `attachment; filename="`+info.Name()+`"`)

		// Serve content (supports Range, HEAD, etc.)
		http.ServeContent(w, r, info.Name(), info.ModTime(), file)
	}
}

func detectMIME(name string) string {
	switch ext := filepath.Ext(name); ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".pdf":
		return "application/pdf"
	case ".txt":
		return "text/plain; charset=utf-8"
	default:
		return "application/octet-stream"
	}
}
