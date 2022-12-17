package command

import (
	"github.com/rocketblend/rocketblend-collector/pkg/cmd/cli/config"
	"github.com/rocketblend/rocketblend-collector/pkg/collector"
	"github.com/rocketblend/scribble"
	"github.com/spf13/cobra"
)

type Service struct {
	config    *config.Config
	driver    *scribble.Driver
	collector *collector.Collector
}

func NewService(config *config.Config, driver *scribble.Driver, collector *collector.Collector) *Service {
	return &Service{
		config:    config,
		driver:    driver,
		collector: collector,
	}
}

func NewCommand(srv *Service) *cobra.Command {
	c := &cobra.Command{
		Use:   "registry-cli",
		Short: "CLI tool for generating Blender build configurations",
		Long:  ``,
	}
	c.SetVersionTemplate("{{.Version}}\n")

	pullCmd := NewPullCommand(srv)

	c.AddCommand(pullCmd)

	return c
}
