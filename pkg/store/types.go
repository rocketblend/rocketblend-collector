package store

import (
	"time"

	"github.com/rocketblend/rocketblend/pkg/driver/runtime"
	"github.com/rocketblend/rocketblend/pkg/semver"
)

type (
	Source struct {
		Platform    runtime.Platform
		FileName    string
		DownloadUrl string
		CreatedAt   time.Time
	}

	Build struct {
		Name      string
		Version   *semver.Version
		Sources   []Source
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
