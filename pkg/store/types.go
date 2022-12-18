package store

import "time"

type (
	Source struct {
		Platform    string
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
