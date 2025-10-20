package middleware

import (
	"log"
	"net/http"
	"time"
)

// responseWriter captures status and body size.
type responseWriter struct {
	http.ResponseWriter
	status int
	bytes  int
}

// Logger wraps any http.Handler and logs each request.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Capture response status via a tiny wrapper
		wrapper := &responseWriter{ResponseWriter: w, status: 200}
		next.ServeHTTP(wrapper, r)

		latency := time.Since(start)
		log.Printf("method=%s url=%s status=%d bytes=%d latency=%.2fms",
			r.Method, r.URL.String(), wrapper.status, wrapper.bytes,
			float64(latency.Microseconds())/1000)
	})
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.bytes += n
	return n, err
}
