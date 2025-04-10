package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"ticket-service/config"
	"ticket-service/controller/handlers"
	"ticket-service/controller/middlewares"
)

type Server struct {
	handlers *handlers.Handlers
	cnf      *config.Config
	Wg       sync.WaitGroup
}

func NewServer(cnf *config.Config, handlers *handlers.Handlers) *Server {
	server := &Server{
		cnf:      cnf,
		handlers: handlers,
	}
	return server
}
func (server *Server) Run() {
	server.Start()
}

func (server *Server) Start() {
	manager := middlewares.NewManager()

	mux := http.NewServeMux()
	//swagger.SetupSwagger(mux, manager)
	server.initRoutes(mux, manager)

	handler := middlewares.EnableCors(mux)

	server.Wg.Add(1)

	go func() {
		defer server.Wg.Done()

		addr := fmt.Sprintf(":%d", server.cnf.HttpPort)

		slog.Info(fmt.Sprintf("Listening at %s", addr))

		if err := http.ListenAndServe(addr, handler); err != nil {
			slog.Error(err.Error())
		}
	}()

}
