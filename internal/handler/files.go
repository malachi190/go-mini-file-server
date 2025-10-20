package handler

import (
	"encoding/json"
	"net/http"

	"github.com/malachi190/go-file-server/internal/storage"
)

// List returns an http.HandlerFunc for GET /list
func List(uploadDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return 
		}

		files, err := storage.List(uploadDir)

		if err != nil {
			http.Error(w, "unable to scan uploads", http.StatusInternalServerError)
			return 
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	}
}