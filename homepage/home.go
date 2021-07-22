package homepage

import (
	"log"
	"net/http"
	"time"
)

const message = "hello"

type Handlers struct {
	log *log.Logger
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		starttime := time.Now()
		defer h.log.Printf("request processed in %s\n", time.Now().Sub(starttime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{log: logger}
}
