package handler

import (
	"log"
	"net/http"

	"github.com/AkankshaNichrelay/Auth-Backend/internal/auth"

	"github.com/go-chi/chi"
)

// Handler used for handling HTTP server requests
type Handler struct {
	Router        *chi.Mux
	log           *log.Logger
	Authenticator *auth.Auth
}

// New returns a new Handler instance
func New(lg *log.Logger, auth *auth.Auth) *Handler {
	mux := chi.NewRouter()

	h := Handler{
		log:           lg,
		Router:        mux,
		Authenticator: auth,
	}

	mux.Get("/", h.getHome)
	mux.Post("/register", h.registerUser)
	mux.Post("/login", h.loginUser)
	mux.Post("/logout", h.logoutUser)

	return &h
}

// getHome returns default landing page
func (h *Handler) getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}
