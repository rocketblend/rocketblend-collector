package collector

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/corpix/uarand"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"github.com/rocketblend/rocketblend-collector/pkg/collection"
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

func DefaultConfig() *Config {
	// TODO: Handle missing config stuff correctly.
	return &Config{
		Proxy:           fmt.Sprintf("https://%s:%s@%s", os.Getenv("PROXY_USER"), os.Getenv("PROXY_PASS"), os.Getenv("PROXY_DOMAIN")),
		UserAgent:       uarand.GetRandom(),
		Parallelism:     2,
		RandomDelay:     time.Second * 5,
		OldestSupported: 2.79,
	}
}

func New(config *Config) *Collector {
	return &Collector{
		conf: config,
	}
}

func (c *Collector) GetStableCollection() *collection.Collection {
	builds := collection.New()

	// TODO: Move collector setup to a separate function/service.
	col := colly.NewCollector(
		colly.AllowedDomains("download.blender.org"),
		colly.UserAgent(c.conf.UserAgent),
		colly.MaxDepth(2),
		colly.Async(true),
	)

	// Proxy
	if c.conf.Proxy != "" {
		rp, err := proxy.RoundRobinProxySwitcher(c.conf.Proxy)
		if err != nil {
			log.Fatal(err)
		}
		col.SetProxyFunc(rp)
	}

	col.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if c.isValidCrawlLink(link) {
			if e.Request.Depth == 1 {
				e.Request.Visit(e.Request.AbsoluteURL(link))
			} else {
				builds.Add(&collection.Build{
					Name:    strings.TrimSuffix(link, filepath.Ext(link)),
					Version: FindVerisonNumberStr(link),
					Tag:     "stable",
					Sources: []collection.Source{
						{
							Platform:    ParsePlatform(link),
							FileName:    link,
							DownloadUrl: e.Request.AbsoluteURL(link),
							CreatedAt:   time.Now(),
						},
					},
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
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

	col.Visit(ReleaseUrl)
	col.Wait()

	return builds
}

func (c *Collector) isValidCrawlLink(url string) bool {
	var expression = "(blender(-)?[0-9]+([.][0-9]+))"
	isValidReleaseName, _ := regexp.MatchString(expression, strings.ToLower(url))

	if isValidReleaseName {
		if version, err := ParseMajorMinorVersionNumber(url); err == nil {
			if version >= c.conf.OldestSupported {
				return true
			}
		}
	}

	return false
}
