package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

type Server struct {   // Dependency
	cnf				*config.Config
	productHandler 	*product.Handler
	userHandler 	*user.Handler
}  // Call all receiver functions

func NewServer(   // Inject as Dependencies
	cnf *config.Config,
	productHandler *product.Handler,     
	userHandler *user.Handler,
) *Server{
	return &Server{
		cnf:			cnf,
		productHandler: productHandler,
		userHandler: 	userHandler,
	}
}

func (server *Server) Start() {    // Receiver Function
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,   // First logger, cors, then preflight run
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	// All routes register
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}