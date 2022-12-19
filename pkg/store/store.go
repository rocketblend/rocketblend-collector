package store

import (
	"sync"
)

type Store struct {
	mu     sync.Mutex
	name   string
	builds map[string]Build
}

func New(name string) *Store {
	return &Store{
		name:   name,
		builds: make(map[string]Build),
	}
}

func (s *Store) GetName() string {
	return s.name
}

func (s *Store) Add(build *Build) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	existing, ok := s.builds[build.Version]
	if ok {
		existing.Sources = append(existing.Sources, build.Sources...)
		existing.UpdatedAt = build.UpdatedAt
		s.builds[build.Version] = existing
		return nil
	}

	s.builds[build.Version] = *build
	return nil
}

func (s *Store) GetAll() map[string]Build {
	return s.builds
}
