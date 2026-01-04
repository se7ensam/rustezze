package jobs

import (
	"errors"
	"sync"
	"fmt"
	"time"
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

func (s *MemoryStore) Dequeue() (*Job, error) {
    fmt.Println("[Store] Dequeue: Waiting for lock to find work...")

    s.mu.Lock()
    defer func() {
        fmt.Println("[Store] Dequeue: Unlocking.")
        s.mu.Unlock()
    }()

    // Iterate through the map to find a 'CREATED' job
    for _, job := range s.data {
        if job.Status == JobStatusCreated {
            fmt.Printf("[Store] Dequeue: Found job %s. Locking it for processing.\n", job.ID)
            
            // Atomic Update: Grab it so no one else can
            job.Status = JobStatusProcessing
            job.UpdatedAt = time.Now()
            
            return job, nil
        }
    }

    // If loop finishes without returning, no work was found
    return nil, nil
}