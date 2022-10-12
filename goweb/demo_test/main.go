package main

import (
	"goweb"
	"log"
	"net/http"
	"time"
)

/*import (
	"goweb"
	"net/http"
)

func main() {
	r := goweb.New()
	r.GET("/index", func(c *goweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello goweb</h1>")
	})
	/*r.GET("/hello", func(c *goweb.Context) {
		// expect /hello?name=gowebktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *goweb.Context) {
		c.JSON(http.StatusOK, goweb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})*/

	/*v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *goweb.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *goweb.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *goweb.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *goweb.Context) {
			c.JSON(http.StatusOK, goweb.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}



	r.Run(":9999")
}*/


func onlyForV2() goweb.HandlerFunc {
	return func(c *goweb.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	//r := goweb.New()
	//r.Use(goweb.Logger()) // global midlleware
	//r.GET("/ljw", func(c *goweb.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//})
	//
	//v2 := r.Group("/v2")
	//v2.Use(onlyForV2()) // v2 group middleware
	//{
	//	v2.GET("/hello/:name", func(c *goweb.Context) {
	//		// expect /hello/geektutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//	})
	//}
	r := goweb.Default()
	r.GET("/panic", func(c *goweb.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})


	r.Run(":9999")
}