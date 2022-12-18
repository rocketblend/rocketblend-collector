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
			collection := srv.collector.GetStableCollection()

			for _, build := range collection.GetAll() {
				srv.driver.Write("builds", build.Version, build)
			}
		},
	}

	return c
}
