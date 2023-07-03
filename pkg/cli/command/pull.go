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
		Long: `Generate build package collections specified by collector.yaml by web-scraping the Blender
build server for the current stable builds.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			fmt.Println("Pulling builds...")

			store, err := srv.collector.CollectStable()
			if err != nil {
				return err
			}

			for _, conf := range srv.config.Collections {
				fmt.Println("Saving '" + conf.Name + "' collection...")

				c := collection.New(srv.config.Library, srv.config.OutputDir, conf.Name, conf.Addons, conf.Platforms, conf.Args, store)

				if err := c.Save(wd); err != nil {
					return err
				}
			}

			return nil
		},
	}

	return c
}
