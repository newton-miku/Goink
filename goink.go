package goink

import (
	"net/http"
)

// 声明HandlerFunc 函数及其参数
type HandlerFunc func(*Context)

type Engine struct {
	// 路由
	router *Router
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handler(c)
}

func (e *Engine) AddRoute(method, path string, handler HandlerFunc) {
	// 调用Router中的addRouter方法
	e.router.addRouter(method, path, handler)
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
	return &Engine{
		router: newRouter(),
	}
}
