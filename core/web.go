package core

import "github.com/gin-gonic/gin"

func NewWeb() *gin.Engine {
	return gin.Default()
}
