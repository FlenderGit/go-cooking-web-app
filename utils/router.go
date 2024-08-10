package utils

import (
	"net/http"
)

type Router struct {
	router *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		router: http.NewServeMux(),
	}
}

func (c *Router) Get(pattern string, handler http.HandlerFunc, middlewares ...Middleware) {
	c.router.HandleFunc(pattern, Chain(handler, middlewares...))
}

func (c *Router) Post(pattern string, handler http.HandlerFunc, middlewares ...Middleware) {
	c.router.HandleFunc(pattern, Chain(handler, middlewares...))
}

func (c *Router) HandleFunc(pattern string, handler http.HandlerFunc, middlewares ...Middleware) {
	c.router.HandleFunc(pattern, Chain(handler, middlewares...))
}

func (c *Router) MountOnServe(mux *http.ServeMux, prefix string) {
	mux.Handle(prefix+"/", http.StripPrefix(prefix, c.router))
}

func (c *Router) MountOnRouter(mux *Router, prefix string) {
	mux.router.Handle(prefix+"/", http.StripPrefix(prefix, c.router))
}
