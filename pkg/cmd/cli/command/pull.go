package command

import (
	"github.com/rocketblend/rocketblend-collector/pkg/collector"
	"github.com/spf13/cobra"
)

func NewPullCommand(srv *Service) *cobra.Command {
	c := &cobra.Command{
		Use:   "pull",
		Short: "Pulls release builds into a local json db",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			collection := srv.collector.GetStableCollection()

			for _, build := range collection.FilterSourcesByPlatform(collector.Platforms[:]) {
				srv.driver.Write("builds", build.Version, build)
			}
		},
	}

	return c
}
