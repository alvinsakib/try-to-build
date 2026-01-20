package product

import (
	"ecommerce/repo"
	middleware "ecommerce/rest/middlewares" // rename
)

type Handler struct {
	middlewares *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middleware.Middlewares,
	productRepo repo.ProductRepo,
) *Handler {
	return &Handler{
		middlewares: middlewares, // Inject
		productRepo: productRepo,
	}
}
