package provider

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) RegisterHandlers(r *gin.Engine) {

	v1 := r.Group("/")
	{

		// Register your endpoints here with the handler
		// fs = First Satellite, ss = Second Satellite, ts = Third Satellite

		v1.POST("/topsecret", dispatchHandler(server.TopSecretHandler.Handle))
		v1.POST("/topsecret_split/:fs_name/:ss_name/:ts_name", dispatchHandler(server.TopSecretSplitHandler.HandlePost))
		v1.GET("/topsecret_split", dispatchHandler(server.TopSecretSplitHandler.HandleGet))

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
