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
func (h *JobHandler) Poll(c *gin.Context) {
    // 1. Ask the Service for work
    job, err := h.Service.PollJob()
    
    // 2. Handle System Errors (Database crashed, etc.)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
        return
    }

    // 3. Handle "No Work" (Success, but empty)
    if job == nil {
        c.Status(http.StatusNoContent) // 204 No Content
        return
    }

    // 4. Handle "Here is a Job"
    c.JSON(http.StatusOK, job)
}