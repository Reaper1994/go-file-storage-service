package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Reaper1994/go-file-storage-service/internal/handlers"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	return &Server{
		Router: mux.NewRouter(),
	}
}

func (s *Server) ConfigureRouter() {
	fileHandler := handlers.NewFileHandler()
	s.Router.HandleFunc("/upload", fileHandler.UploadHandler).Methods("POST")
}

func (s *Server) Start(port string) {
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, s.Router))
}