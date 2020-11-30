package server

import (
	"github.com/go-chi/chi"
	api "go-hansolo/pkg/server/handler"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := chi.NewRouter()

	r.Mount("/api", api.New())
	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

// Start the server.
func (serv *Server) Start() {
	log.Printf("Server running on %s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
