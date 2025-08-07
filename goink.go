package Goink

import (
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (e *Engine) AddRoute(method, path string, handler HandlerFunc) {
	key := method + "-" + path
	e.router[key] = handler
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.AddRoute("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.AddRoute("POST", path, handler)
}

func (e *Engine) PUT(path string, handler HandlerFunc) {
	e.AddRoute("PUT", path, handler)
}

func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.AddRoute("DELETE", path, handler)
}

func (e *Engine) PATCH(path string, handler HandlerFunc) {
	e.AddRoute("PATCH", path, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}
