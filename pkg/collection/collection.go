package collection

import (
	"path"

	"github.com/rocketblend/rocketblend-collector/pkg/store"
	"github.com/rocketblend/rocketblend/pkg/core/library"
	"github.com/rocketblend/rocketblend/pkg/core/runtime"
)

type (
	Collection struct {
		reference string
		packages  []string
		platforms []runtime.Platform
		args      string
		store     *store.Store
	}
)

func New(reference string, packges []string, platforms []runtime.Platform, args string, store *store.Store) *Collection {
	return &Collection{
		reference: reference,
		packages:  packges,
		args:      args,
		platforms: platforms,
		store:     store,
	}
}

func (c *Collection) Map() (output map[string]library.Build, err error) {
	output = make(map[string]library.Build)

	for _, build := range c.store.GetAll() {
		sources := []*library.Source{}
		for _, source := range build.Sources {
			if contains(c.platforms, source.Platform) {
				sources = append(sources, &library.Source{
					Platform:   source.Platform,
					Executable: GetExecutablePath(source.FileName, source.Platform),
					URL:        source.DownloadUrl,
				})
			}
		}
		if len(sources) > 0 {
			output[build.Version] = library.Build{
				Reference: c.reference,
				Args:      c.args,
				Packages:  c.packages,
				Source:    sources,
			}
		}
	}

	return output, nil
}

func GetExecutablePath(dir string, platform runtime.Platform) string {
	return path.Join(dir, platform.String())
}

func contains(s []runtime.Platform, e runtime.Platform) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
