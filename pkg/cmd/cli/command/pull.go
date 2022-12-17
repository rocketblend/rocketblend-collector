package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewPullCommand(srv *Service) *cobra.Command {
	c := &cobra.Command{
		Use:   "pull",
		Short: "Pulls release builds into a local json db",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			collection := srv.collector.GetStableCollection()

			for _, build := range collection.Builds {
				fmt.Println(build)
			}
		},
	}

	return c
}
