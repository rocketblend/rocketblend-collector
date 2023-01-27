package cli

import (
	"fmt"

	cc "github.com/ivanpirog/coloredcobra"

	"github.com/rocketblend/rocketblend-collector/pkg/cmd/cli/command"
	"github.com/rocketblend/rocketblend-collector/pkg/cmd/cli/config"
	"github.com/rocketblend/rocketblend-collector/pkg/collector"
)

func Execute() error {
	config, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	c, err := collector.NewConfig(config.Collector.Proxy, config.Collector.Agent, config.Collector.Parallelism, config.Collector.Delay)
	if err != nil {
		return fmt.Errorf("failed to create collector config: %w", err)
	}

	collector := collector.New(c)
	srv := command.NewService(config, collector)

	rootCMD := command.NewCommand(srv)

	// Configure help template colours
	cc.Init(&cc.Config{
		RootCmd:         rootCMD,
		Headings:        cc.Cyan + cc.Bold + cc.Underline,
		Commands:        cc.Bold,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		Aliases:         cc.Bold,
		Example:         cc.Italic,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})

	return rootCMD.Execute()
}
