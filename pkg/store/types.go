package store

import (
	"time"

	"github.com/rocketblend/rocketblend/pkg/core/runtime"
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
		Version   string
		Tag       string
		Sources   []Source
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
