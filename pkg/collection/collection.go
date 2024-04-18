package collection

import (
	"os"
	"path"
	"path/filepath"

	"github.com/rocketblend/rocketblend-collector/pkg/store"
	"github.com/rocketblend/rocketblend/pkg/helpers"
	"github.com/rocketblend/rocketblend/pkg/runtime"
	"github.com/rocketblend/rocketblend/pkg/types"
)

type (
	Collection struct {
		store     *store.Store
		validator types.Validator
		library   string
		outputDir string
		name      string
		platforms []runtime.Platform
		args      string
	}
)

func New(library string, outputDir string, name string, platforms []runtime.Platform, args string, store *store.Store, validator types.Validator) *Collection {
	return &Collection{
		library:   library,
		outputDir: outputDir,
		name:      name,
		args:      args,
		platforms: platforms,
		store:     store,
		validator: validator,
	}
}

func (c *Collection) GetRoute() string {
	return path.Join(c.outputDir, c.name, "builds", c.store.GetName())
}

func (c *Collection) GetReference() string {
	return path.Join(c.library, c.GetRoute())
}

func (c *Collection) Save(path string) error {
	packs, err := c.convert()
	if err != nil {
		return err
	}

	for version, pack := range packs {
		path := filepath.Join(path, c.GetRoute(), version)
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}

		if err := helpers.Save(c.validator, filepath.Join(path, types.PackageFileName), pack); err != nil {
			return err
		}
	}

	return nil
}

func (c *Collection) convert() (output map[string]*types.Package, err error) {
	output = make(map[string]*types.Package)
	for _, b := range c.store.GetAll() {
		sources := make([]*types.Source, 0, len(b.Sources))
		for _, source := range b.Sources {
			if contains(c.platforms, source.Platform) {
				executable, err := getRuntimeExecutable(source.FileName, source.Platform)
				if err != nil {
					return nil, err
				}

				uri, err := types.NewURI(source.DownloadUrl)
				if err != nil {
					return nil, err
				}

				sources = append(sources, &types.Source{
					Resource: executable,
					URI:      uri,
					Platform: types.Platform(source.Platform.String()),
				})
			}
		}

		if len(sources) > 0 {
			output[b.Version.String()] = &types.Package{
				Type:    types.PackageBuild,
				Version: b.Version,
				Sources: sources,
			}
		}
	}

	return output, nil
}
