package valorant

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/horrorsaur/LAVT/internal/templates/partials"
)

// GET requests are always first handled by the assets FS.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/session", templ.Handler(partials.GetSessionInfo()).ServeHTTP)

	return r
}
