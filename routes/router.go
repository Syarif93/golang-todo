package routes

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	ApiRoutes(router)

	return router
}
