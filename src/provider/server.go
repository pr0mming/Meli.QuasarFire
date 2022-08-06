package provider

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (server *Server) RegisterHandlers(g *gin.Engine) {

	fmt.Println("joa")

	topSecretHdl := server.TopSecretHandler

	g.POST("/topsecret", dispatchHandler(topSecretHdl.Handle))

}

func dispatchHandler(f func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {

	return func(c *gin.Context) {
		response, err := f(c)

		if err != nil {
			c.JSON(404, err)
			return
		}

		c.JSON(200, response)
	}

}
