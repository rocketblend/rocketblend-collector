package collector

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/corpix/uarand"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"github.com/rocketblend/rocketblend-collector/pkg/store"
)

const (
	ReleaseUrl  string = "https://download.blender.org/release/"
	DownloadUrl string = "https://builder.blender.org/download"
)

type Config struct {
	Proxy           string
	UserAgent       string
	Parallelism     int
	RandomDelay     time.Duration
	OldestSupported float32
}

type Collector struct {
	conf *Config
}

func NewConfig(proxyUrl string, agent string, parallelism int, delay string) (*Config, error) {
	duration, err := time.ParseDuration(delay)
	if err != nil {
		return nil, err
	}

	if agent == "random" {
		agent = uarand.GetRandom()
	}

	return &Config{
		Proxy:           proxyUrl,
		UserAgent:       agent,
		Parallelism:     parallelism,
		RandomDelay:     duration,
		OldestSupported: 2.79,
	}, nil
}

func New(config *Config) *Collector {
	return &Collector{
		conf: config,
	}
}

func (c *Collector) CollectStable() (*store.Store, error) {
	builds := store.New("stable")

	col := colly.NewCollector(
		colly.AllowedDomains("download.blender.org"),
		colly.UserAgent(c.conf.UserAgent),
		colly.MaxDepth(2),
		colly.Async(true),
	)

	// Proxy
	if c.conf.Proxy != "" {
		fmt.Printf("Using proxy: %s\n", censorText(c.conf.Proxy, "#", 20))
		rp, err := proxy.RoundRobinProxySwitcher(c.conf.Proxy)
		if err != nil {
			return nil, err
		}

		col.SetProxyFunc(rp)
	}

	col.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if c.isValidCrawlLink(link) {
			if e.Request.Depth == 1 {
				e.Request.Visit(e.Request.AbsoluteURL(link))
			} else {
				now := time.Now()
				version, err := parseVersionNumber(link)
				if version != nil {
					err = builds.Add(&store.Build{
						Name:    strings.TrimSuffix(link, filepath.Ext(link)),
						Version: version,
						Sources: []store.Source{
							{
								Platform:    parsePlatform(link),
								FileName:    link,
								DownloadUrl: e.Request.AbsoluteURL(link),
								CreatedAt:   now,
							},
						},
						CreatedAt: now,
						UpdatedAt: now,
					})
				}

				if err != nil {
					fmt.Println("Failed to add build to collection:", err)
				}
			}
		}
	})

	// Set max Parallelism and introduce a Random Delay
	col.Limit(&colly.LimitRule{
		Parallelism: c.conf.Parallelism,
		RandomDelay: c.conf.RandomDelay,
	})

	// Set error handler
	col.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := col.Visit(ReleaseUrl)
	if err != nil {
		return nil, err
	}

	col.Wait()

	return builds, nil
}

func (c *Collector) isValidCrawlLink(url string) bool {
	var expression = "(blender(-)?[0-9]+([.][0-9]+))"
	isValidReleaseName, _ := regexp.MatchString(expression, strings.ToLower(url))

	if isValidReleaseName {
		if version, err := parseMajorMinorVersionNumber(url); err == nil {
			if version >= c.conf.OldestSupported {
				return true
			}
		}
	}

	return false
}
