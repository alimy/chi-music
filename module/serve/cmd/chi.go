package cmd

import (
	"github.com/alimy/chi-music/module/serve/info"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/unisx/logus"

	mirE "github.com/alimy/mir/module/chi"
)

// newChi return a new chi.Router instance
func newChi() chi.Router {
	// initial chi
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Register API
	entries := info.MirEntries()
	if err := mirE.Register(r, entries...); err != nil {
		logus.F("mir register", err)
	}
	return r
}
