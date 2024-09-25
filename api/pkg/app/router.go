package app

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter(store Store) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", pingHandler)

	v1 := r.Group("/api/v1")
	v1.POST("/dashes", createDashHandler(store))
	v1.GET("/dashes", listDashesHandler(store))

	return r
}
