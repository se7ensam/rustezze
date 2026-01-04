package jobs

import (
	"errors"
	"sync"
	"fmt"
)

type MemoryStore struct {
	data map[string]*Job
	mu   sync.RWMutex 
}


func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]*Job),
	}
}


func (s *MemoryStore) Save(job *Job) {
	fmt.Printf("[Store] Save: %s is waiting for lock...\n", job.ID)
	
	s.mu.Lock()
	
	fmt.Printf("[Store] Save: %s ACQUIRED lock. Writing data.\n", job.ID)


	defer func() {
		fmt.Printf("[Store] Save: %s is unlocking now.\n", job.ID)
		s.mu.Unlock()
	}()

	s.data[job.ID] = job
}


func (s *MemoryStore) Get(id string) (*Job, error) {
	// Debug log for reading
	fmt.Println("[Store] Get: Waiting for Read-Lock...")
	
	s.mu.RLock()
	// Note: We use an anonymous function here too for consistency
	defer func() {
		fmt.Println("[Store] Get: Releasing Read-Lock.")
		s.mu.RUnlock()
	}()
	
	fmt.Println("[Store] Get: Read-Lock acquired. Reading data.")
	
	job, exists := s.data[id]
	if !exists {
		return nil, errors.New("job not found")
	}
	return job, nil
}