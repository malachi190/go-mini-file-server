# Mini File Server – Learning Project in Go

A tiny, self-contained HTTP file server written in Go for educational purposes.
It exposes a REST-ish API and an optional browser UI to upload, list, download, delete and inspect files while demonstrating idiomatic Go patterns: modular handlers, custom middleware, graceful shutdown, embedded static assets, and structured logging.


# ✨ Features
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
