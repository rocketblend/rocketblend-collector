package cli

import (
	"fmt"

	"github.com/rocketblend/rocketblend-collector/pkg/cmd/cli/command"
	"github.com/rocketblend/rocketblend-collector/pkg/cmd/cli/config"
	"github.com/rocketblend/rocketblend-collector/pkg/collector"
)

func Execute() error {
	config, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	c := collector.DefaultConfig()
	collector := collector.New(c)

	srv := command.NewService(config, collector)

	rootCmd := command.NewCommand(srv)
	return rootCmd.Execute()
}
