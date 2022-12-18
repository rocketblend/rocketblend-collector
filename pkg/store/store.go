package store

import (
	"sync"
)

type Store struct {
	mu     sync.Mutex
	builds map[string]Build
}

func New() *Store {
	return &Store{
		builds: make(map[string]Build),
	}
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

func (s *Store) FilterSourcesByPlatform(platforms []string) map[string]Build {
	builds := make(map[string]Build)
	for _, build := range s.builds {
		var filteredSources []Source
		for _, source := range build.Sources {
			if contains(platforms, source.Platform) {
				filteredSources = append(filteredSources, source)
			}
		}
		if len(filteredSources) > 0 {
			build.Sources = filteredSources
			builds[build.Version] = build
		}
	}
	return builds
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
