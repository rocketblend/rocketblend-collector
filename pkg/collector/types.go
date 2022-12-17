package collector

import "time"

type Build struct {
	Platform    string
	Name        string
	Version     string
	Tag         string
	Hash        string
	DownloadUrl string
	CrawledAt   time.Time
}
