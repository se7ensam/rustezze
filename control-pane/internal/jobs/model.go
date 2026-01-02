package jobs

import(
	"time"
)

type JobStatus string

const (
	JobStatusCreated    JobStatus = "CREATED"
	JobStatusQueued     JobStatus = "QUEUED"
	JobStatusProcessing JobStatus = "PROCESSING"
	JobStatusCompleted  JobStatus = "COMPLETED"
	JobStatusFailed     JobStatus = "FAILED"
)
type Job struct
{
	ID        string    `json:"id"`
	Status    JobStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	InputFileURL  string `json:"input_file_url,omitempty"`
	OutputFileURL string `json:"output_file_url,omitempty"`
}

func NewJob(id string) *Job{
	now := time.Now()
	return &Job{
		ID: id,
		Status: JobStatusCreated,
		CreatedAt: now,
		UpdatedAt: now,
	}
}