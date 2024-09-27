package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/stanislavCasciuc/social/internal/store"
)

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string `yaml:"addr"`
	maxOpenConns int    `yaml:"max_open_conns"`
	maxIdleConns int    `yaml:"max_idle_conns"`
	maxIdleTime  string `yaml:"max_idle_time"`
}
type application struct {
	config config
	store  *store.Storage
}

func (a *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", a.heathCheckHandler)

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", a.createPostHandler)
			r.Route("/{postID}", func(r chi.Router) {
				r.Use(a.postsContextMiddleware)
				r.Get("/", a.getPostHandler)
				r.Patch("/", a.updatePostHandler)
				r.Delete("/", a.deletePostHandler)
			})
		})
		r.Route("/users", func(r chi.Router) {
			r.Route("/{userID}", func(r chi.Router) {
				r.Use(a.userContextMiddleware)
				r.Get("/", a.getUserHandler)

				r.Put("/follow", a.followUserHandler)
				r.Put("/unfollow", a.unfollowUserHandler)
			})
		})
	})

	return r
}

func (a *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", a.config.addr)

	return srv.ListenAndServe()
}
