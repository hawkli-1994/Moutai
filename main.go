package main

import "net/http"
import "./wine"

func main() {
	r := wine.New()
	r.GET("/index", func(c *wine.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *wine.Context) {
			c.HTML(http.StatusOK, "<h1>Hello wine</h1>")
		})

		v1.GET("/hello", func(c *wine.Context) {
			// expect /hello?name=lkrMacBookAir
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *wine.Context) {
			// expect /hello/lkrMacBookAir
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *wine.Context) {
			c.JSON(http.StatusOK, wine.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
