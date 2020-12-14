package wine

import (
	"debug/macho"
	//"fmt"
	"net/http"
	//"reflect"
)

type HandlerFunc func(*Context)

//type Route struct {
//	method string
//	pattern string
//	handler HandlerFunc
//}

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) Register(method string, pattern string, handler HandlerFunc) {
	//route := genRouteKey(method, pattern)
	e.router.addRoute(method, pattern, handler)
}



func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.Register("GET", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	c := newContext(w, req)
	e.router.handle(c)
	//route := genRouteKey(req.Method, req.URL.Path)
	//handler, ok := e.routers[route]
	//if ok {
	//	handler(w, req)
	//} else {
	//	fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	//}
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}