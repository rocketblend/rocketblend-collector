package collection

import (
	"os"
	"path"
	"path/filepath"

	"github.com/rocketblend/rocketblend-collector/pkg/store"
	"github.com/rocketblend/rocketblend/pkg/downloader"
	"github.com/rocketblend/rocketblend/pkg/driver/reference"
	"github.com/rocketblend/rocketblend/pkg/driver/rocketpack"
	"github.com/rocketblend/rocketblend/pkg/driver/runtime"
)

type (
	Collection struct {
		store     *store.Store
		library   string
		outputDir string
		name      string
		addons    []reference.Reference
		platforms []runtime.Platform
		args      string
	}
)

func New(library string, outputDir string, name string, addons []reference.Reference, platforms []runtime.Platform, args string, store *store.Store) *Collection {
	return &Collection{
		library:   library,
		outputDir: outputDir,
		name:      name,
		addons:    addons,
		args:      args,
		platforms: platforms,
		store:     store,
	}
}

func (c *Collection) GetRoute() string {
	return path.Join(c.outputDir, c.name, "builds", c.store.GetName())
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

		if err := rocketpack.Save(filepath.Join(buildPath, rocketpack.FileName), b); err != nil {
			return err
		}
	}

	return nil
}

func (c *Collection) convert() (output map[string]*rocketpack.RocketPack, err error) {
	output = make(map[string]*rocketpack.RocketPack)
	for _, b := range c.store.GetAll() {
		sources := make(map[runtime.Platform]*rocketpack.Source)
		for _, source := range b.Sources {
			if contains(c.platforms, source.Platform) {
				executable, err := getRuntimeExecutable(source.FileName, source.Platform)
				if err != nil {
					return nil, err
				}

				uri, err := downloader.NewURI(source.DownloadUrl)
				if err != nil {
					return nil, err
				}

				sources[source.Platform] = &rocketpack.Source{
					Resource: executable,
					URI:      uri,
				}
			}
		}

		if len(sources) > 0 {
			output[b.Version.String()] = &rocketpack.RocketPack{
				Build: &rocketpack.Build{
					Version: b.Version,
					Args:    c.args,
					Addons:  c.addons,
					Sources: sources,
				},
			}
		}
	}

	return output, nil
}
