package storage

import (
	"sync"
)

type ResultStore struct {
	mu      sync.Mutex
	results map[string]float64
}

func NewResultStore() *ResultStore {
	return &ResultStore{
		results: make(map[string]float64),
	}
}

func (s *ResultStore) Save(key string, value float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.results[key] = value
}

func (s *ResultStore) Get(key string) (float64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.results[key]
	return val, ok
}

func (s *ResultStore) GetAll() map[string]float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	copy := make(map[string]float64, len(s.results))
	for k, v := range s.results {
		copy[k] = v
	}
	return copy
}
