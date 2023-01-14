package collection

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"

	"github.com/rocketblend/rocketblend-collector/pkg/store"
	"github.com/rocketblend/rocketblend/pkg/core/build"
	"github.com/rocketblend/rocketblend/pkg/core/runtime"
)

type (
	Collection struct {
		library   string
		name      string
		addons    []string
		platforms []runtime.Platform
		args      string
		store     *store.Store
	}
)

func New(library string, name string, addons []string, platforms []runtime.Platform, args string, store *store.Store) *Collection {
	return &Collection{
		library:   library,
		name:      name,
		addons:    addons,
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

	for version, b := range builds {
		buildPath := filepath.Join(path, c.GetRoute(), version)
		if err := os.MkdirAll(buildPath, 0755); err != nil {
			return err
		}

		buildJSON, err := json.Marshal(b)
		if err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join(buildPath, build.BuildFile), buildJSON, 0644); err != nil {
			return err
		}
	}

	return nil
}

func (c *Collection) convert() (output map[string]build.Build, err error) {
	output = make(map[string]build.Build)

	for _, b := range c.store.GetAll() {
		sources := []*build.Source{}
		for _, source := range b.Sources {
			if contains(c.platforms, source.Platform) {
				executable, err := getRuntimeExecutable(source.FileName, source.Platform)
				if err != nil {
					return nil, err
				}
				sources = append(sources, &build.Source{
					Platform:   source.Platform,
					Executable: executable,
					URL:        source.DownloadUrl,
				})
			}
		}
		if len(sources) > 0 {
			output[b.Version.String()] = build.Build{
				BlenderVersion: b.Version,
				Args:           c.args,
				Addons:         c.addons,
				Source:         sources,
			}
		}
	}

	return output, nil
}
