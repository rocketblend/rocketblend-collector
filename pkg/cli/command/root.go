package command

import (
	"github.com/rocketblend/rocketblend-collector/pkg/cli/config"
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
		Short: "Collector is a package generator tool for RocketBlend.",
		Long: `Collector is a command-line tool for generating package configurations quickly
and easily for use with RocketBlend.

Documentation is available at https://docs.rocketblend.io/v/collector/`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	c.SetVersionTemplate("{{.Version}}\n")

	c.AddCommand(NewPullCommand(srv))

	return c
}
