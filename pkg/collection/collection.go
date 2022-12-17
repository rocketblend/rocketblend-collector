package collection

type Collection struct {
	Builds []Build
}

func New() *Collection {
	return &Collection{}
}

func (c *Collection) Add(build *Build) {
	existing := c.Get(build.Version)
	if existing != nil {
		existing.Sources = append(existing.Sources, build.Sources...)
		existing.UpdatedAt = build.UpdatedAt
		return
	}

	c.Builds = append(c.Builds, *build)
}

func (c *Collection) Get(version string) *Build {
	for _, build := range c.Builds {
		if build.Version == version {
			return &build
		}
	}
	return nil
}
