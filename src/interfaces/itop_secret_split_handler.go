package interfaces

import "github.com/gin-gonic/gin"

type ITopSecretSplitHandler interface {
	HandleGet(c *gin.Context) (interface{}, error)
	HandlePost(c *gin.Context) (interface{}, error)
}
