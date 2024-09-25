package app

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter(store *Store) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/ping", pingHandler)
	v1.POST("/dashes", createDashHandler(store))
	v1.GET("/dashes", listDashesHandler(store))

	return r
}
