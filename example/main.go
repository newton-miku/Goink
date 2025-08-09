package main

import (
	"net/http"

	goink "github.com/newton-miku/Goink"
)

func main() {
	engine := goink.New()
	engine.GET("/", func(ctx *goink.Context) {
		ctx.String(http.StatusOK, "Hello World")
	})
	engine.GET("/hello", func(ctx *goink.Context) {
		ctx.Stringf(http.StatusOK, "Hello %s , you are accessing %s\n", ctx.Query("name"), ctx.Path)
	})

	engine.POST("/login", func(ctx *goink.Context) {
		ctx.JSON(http.StatusOK, goink.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	engine.Run(":9999")
}
