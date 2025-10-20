package handler

import (
	"net/http"

	"github.com/malachi190/go-file-server/internal/storage"
)

// Delete returns http.HandlerFunc for DELETE /delete?name=filename
func Delete(uploadDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		filename := r.URL.Query().Get("name")

		if filename == "" {
			http.Error(w, "missing file name", http.StatusBadRequest)
			return
		}

		err := storage.Delete(uploadDir, filename)

		if err != nil {
			switch err {
			case storage.ErrNotFound:
				http.Error(w, "file not found", http.StatusNotFound)
			default:
				http.Error(w, "failed to delete", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
