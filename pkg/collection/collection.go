package collection

import (
	"sync"
)

type Collection struct {
	mu     sync.Mutex
	builds map[string]Build
}

func New() *Collection {
	return &Collection{
		builds: make(map[string]Build),
	}
}

func (c *Collection) Add(build *Build) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	existing, ok := c.builds[build.Version]
	if ok {
		existing.Sources = append(existing.Sources, build.Sources...)
		existing.UpdatedAt = build.UpdatedAt
		c.builds[build.Version] = existing
		return nil
	}

	c.builds[build.Version] = *build
	return nil
}

func (c *Collection) GetAll() map[string]Build {
	return c.builds
}

func (c *Collection) FilterByPlatform(platforms []string) map[string]Build {
	filteredBuilds := make(map[string]Build)
	for _, build := range c.builds {
		var filteredSources []Source
		for _, source := range build.Sources {
			if contains(platforms, source.Platform) {
				filteredSources = append(filteredSources, source)
			}
		}
		build.Sources = filteredSources
		filteredBuilds[build.Version] = build
	}
	return filteredBuilds
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
