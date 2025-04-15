package web

import (
	"net/http"
	"ticket-service/controller/middlewares"
)

func (server *Server) initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /hello",
		manager.With(
			http.HandlerFunc(server.handlers.Hello),
		),
	)

	mux.Handle(
		"GET /get-ticket",
		manager.With(
			http.HandlerFunc(server.handlers.GetTicket),
			server.middleware.AuthenticateTicket,
		),
	)

	mux.Handle(
		"GET /get-tickets",
		manager.With(
			http.HandlerFunc(server.handlers.GetTickets),
		),
	)

	mux.Handle(
		"DELETE /cancle-processing",
		manager.With(
			http.HandlerFunc(server.handlers.CancleProcessing),
		),
	)

	mux.Handle(
		"POST /create-ticket",
		manager.With(
			http.HandlerFunc(server.handlers.Createtickets),
		),
	)

	mux.Handle(
		"POST /confirm-ticket",
		manager.With(
			http.HandlerFunc(server.handlers.ConfirmTicket),
			server.middleware.AuthenticateTicket,
		),
	)
}
