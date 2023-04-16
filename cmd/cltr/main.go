package main

import (
	"fmt"
	"os"

	"github.com/rocketblend/rocketblend-collector/pkg/cltr"
)

func main() {
	if err := cltr.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
