package main

import "net/http"
import "./wine"

func main() {
	r := wine.New()
	r.GET("/", func(c *wine.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Moutai</h1>")
	})

	r.GET("/hello", func(c *wine.Context) {
		// expect /hello?name=lkrMacBook
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *wine.Context) {
		// expect /hello/lkrMacBook
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *wine.Context) {
		c.JSON(http.StatusOK, wine.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}