package wine

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func genRouteKey(method string, pattern string) string {
	return method + "-" + pattern
}

func newRouter() *router  {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %s - %s", method, pattern)
	key := genRouteKey(method, pattern)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := genRouteKey(c.Method, c.Path)
	handler, ok := r.handlers[key]
	if ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	}
}