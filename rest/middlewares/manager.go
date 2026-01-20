package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware // [CorsWithflight, Logger, Hudai]
}

func NewManager() *Manager {

}

func (mngr *Manager) Use(middlewars ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewars...)
}

func (mngr *Manager) With(handler http.Handler, middleware ...Middleware) http {
	h := handler

	for _, middleware := range middlewares {
		h = middleware(h)
	}
}
