package collection

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/rocketblend/rocketblend-collector/pkg/store"
	"github.com/rocketblend/rocketblend/pkg/core/library"
	"github.com/rocketblend/rocketblend/pkg/core/runtime"
)

type (
	Collection struct {
		library   string
		name      string
		packages  []string
		platforms []runtime.Platform
		args      string
		store     *store.Store
	}
)

func New(library string, name string, packges []string, platforms []runtime.Platform, args string, store *store.Store) *Collection {
	return &Collection{
		library:   library,
		name:      name,
		packages:  packges,
		args:      args,
		platforms: platforms,
		store:     store,
	}
}

func (c *Collection) GetRoute() string {
	return path.Join("builds", c.name, c.store.GetName())
}

func (c *Collection) GetReference() string {
	return path.Join(c.library, c.GetRoute())
}

func (c *Collection) Save(path string) error {
	builds, err := c.convert()
	if err != nil {
		return err
	}

	for version, build := range builds {
		buildPath := filepath.Join(path, c.GetRoute(), version)
		if err := os.MkdirAll(buildPath, 0755); err != nil {
			return err
		}

		buildJSON, err := json.Marshal(build)
		if err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join(buildPath, "build.json"), buildJSON, 0644); err != nil {
			return err
		}
	}

	return nil
}

func (c *Collection) convert() (output map[string]library.Build, err error) {
	output = make(map[string]library.Build)

	for _, build := range c.store.GetAll() {
		sources := []*library.Source{}
		for _, source := range build.Sources {
			if contains(c.platforms, source.Platform) {
				sources = append(sources, &library.Source{
					Platform:   source.Platform,
					Executable: path.Join(trimSuffix(source.FileName), getRuntimeExecutable(source.Platform)),
					URL:        source.DownloadUrl,
				})
			}
		}
		if len(sources) > 0 {
			output[build.Version.String()] = library.Build{
				Reference:      filepath.Join(c.GetReference(), build.Version.String()),
				BlenderVersion: build.Version,
				Args:           c.args,
				Packages:       c.packages,
				Source:         sources,
			}
		}
	}

	return output, nil
}

func trimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
