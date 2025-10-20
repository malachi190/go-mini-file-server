package handler

import (
	"encoding/json"
	"net/http"

	"github.com/malachi190/go-file-server/internal/storage"
)

// Metadata returns http.HandlerFunc for GET /metadata?name=filename
func Metadata(uploadDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "missing query param 'name'", http.StatusBadRequest)
			return
		}

		meta, err := storage.GetMeta(uploadDir, name)
		if err != nil {
			switch err {
			case storage.ErrNotFound:
				http.Error(w, "file not found", http.StatusNotFound)
			default:
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(meta)
	}
}
