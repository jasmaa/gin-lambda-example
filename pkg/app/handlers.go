package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func createDashHandler(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request CreateDashRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dash := Dash{
			DashId:    uuid.New().String(),
			Name:      request.Name,
			Speed:     request.Speed,
			CreatedAt: time.Now(),
		}

		s.CreateDash(dash)

		c.JSON(http.StatusOK, CreateDashResponse{
			Dash: dash,
		})
	}
}

func listDashesHandler(s *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		dashes := s.ListDashes()
		c.JSON(http.StatusOK, ListDashResponse{
			Dashes: dashes,
		})
	}
}
