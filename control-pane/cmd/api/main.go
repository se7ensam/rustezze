package main

import (
	
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"rusteze/control-pane/internal/jobs"
	"github.com/google/uuid"
)

func main(){
	r := gin.Default()


	jobStore := make(map[string]*jobs.Job)
	r.GET("/ping", func (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":"Control pane is online",
		"system": "rusteze",
	})
})
   

	r.POST("/upload",func (c *gin.Context){
		jobID := "job_12345"
		fmt.Println("Received upload request. Creating job:", jobID)
		id := uuid.New().String()

		newJob := jobs.NewJob(id)

		jobStore[id] = newJob

		c.JSON(http.StatusAccepted, newJob)
	})

	fmt.Println("Starting server on port 8080")
	r.Run(":8080")
}