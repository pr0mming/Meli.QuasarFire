package provider

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) RegisterHandlers(r *gin.Engine) {

	v1 := r.Group("/")
	{
		topSecretHdl := server.TopSecretHandler

		// Register your endpoints here with the handler
		v1.POST("/topsecret", dispatchHandler(topSecretHdl.Handle))

	}

}

func dispatchHandler(f func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {

	return func(c *gin.Context) {
		response, err := f(c)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}

}
