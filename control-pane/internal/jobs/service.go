package jobs

import (
	"time"
)


type JobService struct {
	store *MemoryStore
}


func NewService(store *MemoryStore) *JobService {
	return &JobService{
		store: store,
	}
}


func (s *JobService) CreateJob(id string) *Job {
	job := NewJob(id)
	

	s.store.Save(job)
	
	return job
}


func (s *JobService) GetJob(id string) (*Job, error) {
	return s.store.Get(id)
}
func (s *JobService) PollJob() (*Job, error) {
    return s.store.Dequeue()
}


func (s *JobService) UpdateStatus(id string, status JobStatus) error {
	job, err := s.store.Get(id)
	if err != nil {
		return err
	}
	
	job.Status = status
	job.UpdatedAt = time.Now()
	s.store.Save(job)
	
	return nil
}