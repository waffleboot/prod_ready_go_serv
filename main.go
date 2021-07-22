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
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage.HomeHandler)

	srv := server.New(mux, GcukServiceAddr)
	err := srv.ListenAndServeTLS(GcukCertFile, GcukKeyFile)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
