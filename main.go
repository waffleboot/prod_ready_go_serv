package main

import (
	"log"
	"net/http"
	"os"

	"gopherconuk/server"
)

const message = "hello"

var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	srv := server.New(mux, GcukServiceAddr)
	err := srv.ListenAndServeTLS(GcukCertFile, GcukKeyFile)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
