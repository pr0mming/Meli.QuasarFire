package main

import (
	"github.com/gin-gonic/gin"
	"meli.quasarfire/src/injector"
)

func main() {
	g := gin.Default()

	server := injector.Wire()

	server.RegisterHandlers(g)

	g.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
