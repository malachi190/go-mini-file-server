package handler

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/malachi190/go-file-server/internal/storage"
)

type UploadRes struct {
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	Message string `json:"message"`
}

// Upload returns an http.HandlerFunc for POST /upload
func Upload(uploadDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse mutltipart form, max should be 32mb in memory, rest to temp file.
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, "failed to parse form: "+err.Error(), http.StatusBadRequest)
			return
		}

		file, header, err := r.FormFile("file")

		if err != nil {
			http.Error(w, "missing 'file' field"+err.Error(), http.StatusBadRequest)
			return
		}

		defer file.Close()

		// basic filename sanity check
		filename := strings.TrimSpace(filepath.Base(header.Filename))
		if filename == "" || strings.Contains(filename, "..") {
			http.Error(w, "invalid filename", http.StatusBadRequest)
			return
		}

		// save file
		size, err := storage.Save(uploadDir, filename, file)

		if err != nil {
			http.Error(w, "file save failed:"+err.Error(), http.StatusInternalServerError)
			return
		}

		// return response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UploadRes{
			Name:    filename,
			Size:    size,
			Message: "Upload successful",
		})
	}
}
