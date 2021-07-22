package main

import (
	"log"
	"net/http"
	"os"

	"gopherconuk/homepage"
	"gopherconuk/server"
)

var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage.HomeHandler)

	srv := server.New(mux, GcukServiceAddr)

	logger.Println("server starting")
	err := srv.ListenAndServeTLS(GcukCertFile, GcukKeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
