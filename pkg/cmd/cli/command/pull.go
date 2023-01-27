package command

import (
	"fmt"
	"os"

	"github.com/rocketblend/rocketblend-collector/pkg/collection"
	"github.com/spf13/cobra"
)

func NewPullCommand(srv *Service) *cobra.Command {
	c := &cobra.Command{
		Use:   "pull",
		Short: "Generate build package collections for the current stable builds",
		Long: `generate build package collections specified by collector.yaml by web-scraping the Blender
build server for the current stable builds.`,
		Run: func(cmd *cobra.Command, args []string) {
			wd, _ := os.Getwd()

			fmt.Println("pulling builds...")
			store := srv.collector.CollectStable()
			fmt.Println("done pulling builds")

			for _, conf := range *srv.config.Collections {
				fmt.Println("saving collection: " + conf.Name)
				c := collection.New(srv.config.Library, conf.Name, conf.Addons, conf.Platforms, conf.Args, store)
				if err := c.Save(wd); err != nil {
					fmt.Printf("failed to save collection: %s", err)
				} else {
					fmt.Println("done saving collection")
				}
			}
		},
	}

	return c
}
