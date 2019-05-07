package main

import (
	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func route() (c *chi.Mux) {
	c = chi.NewMux()

	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Use(checkCookie)

	notFound := ShowNotFound()
	c.NotFound(notFound)
	c.Get("/static/*", ServeStatic(notFound))

	c.Get("/", ShowMain())
	c.Post("/submit", SubmitPost())
	c.Post("/vote", VotePost())

	return
}
