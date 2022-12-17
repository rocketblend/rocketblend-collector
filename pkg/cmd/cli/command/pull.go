package command

import (
	"github.com/spf13/cobra"
)

func NewPullCommand(srv *Service) *cobra.Command {
	c := &cobra.Command{
		Use:   "pull",
		Short: "Pulls release builds into a local json db",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			builds := srv.collector.GetStableBuilds()

			for _, build := range builds {
				srv.driver.Write("build", build.Hash, build)
			}
		},
	}

	return c
}
