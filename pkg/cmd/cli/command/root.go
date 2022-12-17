package command

import (
	"github.com/rocketblend/rocketblend-collector/pkg/collector"
	"github.com/rocketblend/scribble"
	"github.com/spf13/cobra"
)

type Service struct {
	driver    *scribble.Driver
	collector *collector.Collector
}

func NewService(driver *scribble.Driver, collector *collector.Collector) *Service {
	return &Service{
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
