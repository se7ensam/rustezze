package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	
	"rusteze/control-pane/internal/jobs"
)


type JobHandler struct {
	Service *jobs.JobService
}


func NewJobHandler(s *jobs.JobService) *JobHandler {
	return &JobHandler{
		Service: s,
	}
}

func (h *JobHandler) CreateJob(c *gin.Context) {

	id := uuid.New().String()


	job := h.Service.CreateJob(id)

	
	c.JSON(http.StatusAccepted, job)
}

func (h *JobHandler) GetJob(c *gin.Context) {
	id := c.Param("id")

	job, err := h.Service.GetJob(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	c.JSON(http.StatusOK, job)
}