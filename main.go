package main

import "net/http"
import "./moutai"

func main() {
	r := moutai.Default()

	v1 := r.Group("/v1")
	{


		v1.GET("/hello", func(c *moutai.Context) {
			// expect /hello?name=lkrMacBookAir
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *moutai.Context) {
			// expect /hello/lkrMacBookAir
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *moutai.Context) {
			c.JSON(http.StatusOK, moutai.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
