package cli

import (
	"github.com/rocketblend/rocketblend-collector/pkg/cmd/cli/command"
	"github.com/rocketblend/rocketblend-collector/pkg/collector"
	"github.com/rocketblend/scribble"
)

func Execute() error {
	driver, err := scribble.New("data/", nil)
	if err != nil {
		return nil
	}

	config := collector.DefaultConfig()
	collector := collector.New(config)

	srv := command.NewService(driver, collector)

	rootCmd := command.NewCommand(srv)
	return rootCmd.Execute()
}
