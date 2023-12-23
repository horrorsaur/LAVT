package valorant

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// GET requests are always first handled by the assets FS.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// r.Get("/", templ.Handler())
	// components.Header("Less Annoying Valorant Tracker")).ServeHTTP

	return r
}
