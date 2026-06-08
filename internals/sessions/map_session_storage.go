package sessions

import (
	"os"
	"sync"
)

type MockSessionsStorage struct {
	mu      sync.RWMutex
	storage map[string]string
}

func (s *MockSessionsStorage) Store(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.storage[key] = value

	return nil
}

func (s *MockSessionsStorage) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, exists := s.storage[key]

	if !exists {
		return "", os.ErrNotExist
	} else {
		return value, nil
	}
}
