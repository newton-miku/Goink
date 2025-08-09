package goink

import "net/http"

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) addRouter(method, path string, handler HandlerFunc) {
	key := method + "-" + path
	r.handlers[key] = handler
}

func (r *Router) handler(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND")
	}
}
