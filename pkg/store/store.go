package store

import (
	"fmt"
	"sync"
)

type Store struct {
	mu     sync.Mutex
	name   string
	builds map[string]*Build
}

func New(name string) *Store {
	return &Store{
		name:   name,
		builds: make(map[string]*Build),
	}
}

func (s *Store) GetName() string {
	return s.name
}

func (s *Store) Add(build *Build) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if build.Version == nil {
		return fmt.Errorf("build version cannot be nil")
	}

	version := build.Version.String()
	existing, ok := s.builds[version]
	if ok {
		existing.Sources = append(existing.Sources, build.Sources...)
		existing.UpdatedAt = build.UpdatedAt
		s.builds[version] = existing
		return nil
	}

	s.builds[version] = build
	return nil
}

func (s *Store) GetAll() map[string]*Build {
	return s.builds
}
