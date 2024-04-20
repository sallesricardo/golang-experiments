package handlers

import (
    "github.com/go-chi/chi"
    chimiddle "github.com/go-chi/chi/middleware"
    "github.com/sallesricardo/golang-experiments/basic/api/internal/middleware"
)

func Handlers(r *chi.Mux) {
    // global middleware
    r.Use(chimiddle.StripSlashes)

    r.Router("/account", func (router chi.Router) {
        // middleware for /account
        router.Use(middleware.Authorization)

        router.Get("/coins", GetCoinBalance)
    })
}
