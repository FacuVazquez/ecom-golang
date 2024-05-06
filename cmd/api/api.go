package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/FacuVazquez/ecom-golang/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(router)

	subrouter := http.NewServeMux()
	subrouter.Handle("/api/v1", http.StripPrefix("/api/v1", router))

	log.Println("Listening on ", s.addr)

	return http.ListenAndServe(s.addr, subrouter)
}
