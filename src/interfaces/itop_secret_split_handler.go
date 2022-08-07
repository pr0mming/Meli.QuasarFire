package interfaces

import "github.com/gin-gonic/gin"

type ITopSecretSplitHandler interface {
	Handle(c *gin.Context) (interface{}, error)
}
