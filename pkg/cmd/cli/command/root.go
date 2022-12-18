package command

import (
	"github.com/rocketblend/rocketblend-collector/pkg/cmd/cli/config"
	"github.com/rocketblend/rocketblend-collector/pkg/collector"
	"github.com/spf13/cobra"
)

type Service struct {
	config    *config.Config
	collector *collector.Collector
}

func NewService(config *config.Config, collector *collector.Collector) *Service {
	return &Service{
		config:    config,
		collector: collector,
	}
}

func NewCommand(srv *Service) *cobra.Command {
	c := &cobra.Command{
		Use:   "collector",
		Short: "CLI tool for generating Blender build configurations",
		Long:  ``,
	}
	c.SetVersionTemplate("{{.Version}}\n")

	pullCmd := NewPullCommand(srv)

	c.AddCommand(pullCmd)

	return c
}
