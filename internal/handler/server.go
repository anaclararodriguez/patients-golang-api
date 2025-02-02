package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Port   string
	Router *mux.Router
}

func NewServer(port string) *Server {
	s := &Server{
		Port:   port,
		Router: mux.NewRouter(),
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	api := s.Router.PathPrefix("/").Subrouter()

	api.Handle("/", http.HandlerFunc(RootHandler))
	api.Handle("/patients", http.HandlerFunc(GetPatients)).Methods("GET")
	api.Handle("/patients", http.HandlerFunc(CreatePatient)).Methods("POST")
	api.Handle("/{anything:.*}", http.HandlerFunc(InvalidPathHandler))

}

func (s *Server) Start() {
	log.Printf("Starting patients-golang-api at http://localhost:%s", s.Port)
	if err := http.ListenAndServe(":"+s.Port, s.Router); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
