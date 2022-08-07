package interfaces

import "github.com/gin-gonic/gin"

type ITopSecretHandler interface {
	Handle(c *gin.Context) (interface{}, error)
}
