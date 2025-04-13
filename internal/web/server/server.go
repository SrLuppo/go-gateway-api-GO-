package server

import (
	"context"
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/web/handlers"
	middleware "github.com/devfullcycle/imersao22/go-gateway/internal/web/midddleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	services       *http.Server
	AccountService *service.AccountService
	InvoiceService *service.InvoiceService
	port           string
}

func NewServer(port string, accountService *service.AccountService, invoiceService *service.InvoiceService) *Server {
	return &Server{
		router:         chi.NewRouter(),
		port:           port,
		AccountService: accountService,
		InvoiceService: invoiceService,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.AccountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.InvoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.AccountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		r.Post("/invoices", invoiceHandler.Create)
		r.Get("/invoices", invoiceHandler.ListByAccount)
		r.Get("/invoices/{id}", invoiceHandler.GetByID)
	})

}

func (s *Server) Start() error {
	s.ConfigureRoutes()

	s.services = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.services.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.services.Shutdown(context.Background())
}
