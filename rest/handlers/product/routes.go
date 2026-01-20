package product

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandleFunc(h.GetProducts),
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandleFunc(h.CreateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandleFunc(h.GetProduct),
		),
	)

	mux.Handle(
		"PUT /product/{id}",
		manager.With(
			http.HandleFunc(h.UpdateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandleFunc(h.DeleteProduct),
			h.middlewares.AuthenticateJWT,
		),
	)
}