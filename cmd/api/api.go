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
	router.HandleFunc("POST /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte("User ID: " + userID))
	})

	// This is a subrouter
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(v1)

	server := http.Server{
		Addr:    s.addr,
		Handler: v1,
	}

	log.Println("server has started")

	return server.ListenAndServe()
}
