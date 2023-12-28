package local

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/horrorsaur/LAVT/internal/templates/partials"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

type (
	Router struct{}
)

// GET requests are always first handled by the assets FS.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/session", templ.Handler(partials.GetSessionInfo()).ServeHTTP)
	r.Get("/connect", connectToWS)

	return r
}

func connectToWS(w http.ResponseWriter, r *http.Request) {
}

func addRoute(r *chi.Mux, path string, handler http.HandlerFunc) {
	r.Get(path, handler)
}
