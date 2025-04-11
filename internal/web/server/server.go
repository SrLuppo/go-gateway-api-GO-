package server

import (
	"context"
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	services       *http.Server
	AccountService service.AccountService
	port           string
}

func NewServer(port string, accountService service.AccountService) *Server {
	return &Server{
		router:         chi.NewRouter(),
		port:           port,
		AccountService: accountService,
	}
}

func (s *Server) Start() error {
	accountHandler := handlers.NewAccountHandler(s.AccountService)
	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	s.services = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.services.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.services.Shutdown(context.Background())
}
