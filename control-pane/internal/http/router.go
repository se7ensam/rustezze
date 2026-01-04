package http

import (
	"github.com/gin-gonic/gin"
)


func NewRouter(handler *JobHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/upload", handler.CreateJob)
	r.GET("/jobs/:id", handler.GetJob)
	r.POST("/jobs/poll", handler.Poll)

	return r
}