package interfaces

import "github.com/gin-gonic/gin"

type IServer interface {
	RegisterHandlers(g *gin.Engine)
}
