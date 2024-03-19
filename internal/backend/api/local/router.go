package local

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// todo: maybe swap this out with Go Fiber since I like that better

// GET requests are always first handled by the Wails embedded assets (see main).
// chi.Mux implements the http.Handler interface
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	return r
}
