package config

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

/**
setup route to all app
*/
func SetupRouter() *gin.Engine {
	return gin.Default()
}


