package collection

import "github.com/rocketblend/rocketblend-collector/pkg/store"

type (
	Collection struct {
		name      string
		platforms []string
		store     *store.Store
	}
)

func New(name string, platforms []string, store *store.Store) *Collection {
	return &Collection{
		name:      name,
		platforms: platforms,
		store:     store,
	}
}

func (c *Collection) ToOutput() string {
	return c.name
}
