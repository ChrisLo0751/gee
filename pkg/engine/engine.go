package engine

import (
	"fmt"
	"net/http"
)

const (
	GET  RequetsMethod = "GET"
	POST RequetsMethod = "POST"
)

type RequetsMethod string

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRouter(method RequetsMethod, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	e.router[key] = handler
}

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.addRouter(GET, pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandlerFunc) {
	e.addRouter(POST, pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := fmt.Sprintf("%s-%s", req.Method, req.URL.Path)
	if h, ok := e.router[key]; ok {
		h(w, req)
	} else {
		fmt.Fprintf(w, "404 Not Found")
	}
}
