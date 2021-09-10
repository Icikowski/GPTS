package health

import "sync"

type status struct {
	readiness bool
	mutex     sync.Mutex
}

// GetStatus gets the current status of monitored entity
func (s *status) GetStatus() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.readiness
}

// SetStatus sets the current status of monitored entity
func (s *status) SetStatus(ready bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.readiness = ready
}
