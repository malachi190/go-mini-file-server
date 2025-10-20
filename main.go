package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/malachi190/go-file-server/internal/handler"
	"github.com/malachi190/go-file-server/internal/middleware"
	"github.com/malachi190/go-file-server/internal/storage"
)

//go:embed all:static
var staticFS embed.FS

const uploadDir = "./uploads"

func main() {
	fmt.Println("Welcome")

	if err := storage.CreateDir(uploadDir); err != nil {
		log.Fatalf("create uploads dir: %v", err)
	}

	// create a sub FS that strips the leading "static/" prefix
	staticContent, _ := fs.Sub(staticFS, "static")

	mux := http.NewServeMux()
	// health check
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "pong")
	})
	// redirect "/" -> "/static/index.html"
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusFound))
	// serve everything under /static/ (index.html, future css/js, etc.)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticContent))))

	mux.HandleFunc("/upload", handler.Upload(uploadDir))
	mux.HandleFunc("/files/", handler.Download(uploadDir))
	mux.HandleFunc("/list", handler.List(uploadDir))
	mux.HandleFunc("/delete", handler.Delete(uploadDir))
	mux.HandleFunc("/metadata", handler.Metadata(uploadDir))

	loggedHandler := middleware.Logger(mux)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      loggedHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// graceful shutdown goroutine
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig

		log.Println("shutdown signal received, waiting 30 s max...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
		}
		log.Println("server exited")
	}()

	log.Println("listening on http://localhost:8080")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
