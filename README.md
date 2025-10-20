Mini File Server – Learning Project in Go
A tiny, self-contained HTTP file server written in Go for educational purposes.
It exposes a REST-ish API and an optional browser UI to upload, list, download, delete and inspect files while demonstrating idiomatic Go patterns: modular handlers, custom middleware, graceful shutdown, embedded static assets, and structured logging.


✨ Features
Table
Copy
Endpoint	Method	Description
POST /upload	multipart/form-data	Upload a file (form field file)
GET /files/{filename}	GET	Download a file (inline or attachment)
GET /list	GET	JSON array of all files (name, size, created, SHA-256)
DELETE /delete?name=…	DELETE	Remove a file
GET /metadata?name=…	GET	Detailed metadata incl. MIME type & SHA-256
GET /	GET	Tiny browser UI (upload, table, download, delete)
Additional niceties:
Structured request logger (method, URL, status, bytes, latency)
Graceful shutdown (waits up to 30 s for in-flight requests)
Safe path handling (no directory traversal)
Content-Type detection & Content-Disposition: attachment
Range-request support (http.ServeContent)
Single binary (static HTML/CSS/JS embedded with //go:embed)


🚀 Quick Start
bash
Copy
# clone / cd into repo
go mod tidy
go run .              # default port :8080
# or
PORT=3000 go run .
Open http://localhost:8080 – drag-and-drop files or use curl:
bash
Copy
# upload
curl -F "file=@photo.jpg" http://localhost:8080/upload

# list
curl http://localhost:8080/list | jq .

# download
curl -O http://localhost:8080/files/photo.jpg

# delete
curl -X DELETE "http://localhost:8080/delete?name=photo.jpg"


🧪 Project Anatomy
Copy
mini-file-server/
├── main.go                 # wiring, graceful shutdown, embed static
├── internal/
│   ├── handler/            # upload, download, list, delete, metadata
│   ├── logger/             # request logging middleware
│   └── storage/            # disk I/O helpers + SHA-256
├── static/
│   └── index.html          # minimal web UI (embedded)
├── uploads/                # default storage dir (git-ignored)
└── go.mod

Handler flow:
router → middleware.Logger → handler → storage

🔒 Security Notes (for learning only)
Files are stored as-is under ./uploads – add virus scanning, rate-limiting, UUID filenames, etc. before production.
No HTTPS – put behind a reverse proxy like Nginx to add TLS.

🧠 Learning Exercises To Try
Add authentication (Basic, JWT, sessions).
Store files in S3 instead of local disk.
Add pagination or search to /list.
Compute SHA-256 while streaming upload (save a second pass).
Write unit tests for handlers with httptest.
Add WebSocket progress bar for large uploads.
Containerise with a multi-stage Dockerfile.


📄 Licence
MIT – feel free to copy, hack, teach.