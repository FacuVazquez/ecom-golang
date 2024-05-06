package user

import "net/http"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/login", h.handleLogin)
	router.HandleFunc("/register", h.handleResgister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleResgister(w http.ResponseWriter, r *http.Request) {
}
