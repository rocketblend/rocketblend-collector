package main

import (
	"fmt"
	"os"

	"github.com/rocketblend/rocketblend-collector/pkg/cmd/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}