# RocketBlend Collector

[![Github tag](https://badgen.net/github/tag/rocketblend/rocketblend-collector)](https://github.com/rocketblend/rocketblend-collector/tags)
[![Go Doc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/rocketblend/rocketblend-collector)
[![Go Report Card](https://goreportcard.com/badge/github.com/rocketblend/rocketblend-collector)](https://goreportcard.com/report/github.com/rocketblend/rocketblend-collector)
[![GitHub](https://img.shields.io/github/license/rocketblend/rocketblend-collector)](https://github.com/rocketblend/rocketblend-collector/blob/master/LICENSE)

CLI tool for collecting [Blender](https://www.blender.org/) build information for use with use with [RocketBlend](https://github.com/rocketblend/rocketblend)

## Example config

```yaml
library: github.com/rocketblend/official-library
collector:
    proxy: https://proxy.blender.org
    agent: random
    parallelism: 2
    delay: 2s
collections:
    - collection:
        name: Blender
        platforms:
            - linux
            - windows
            - macos
    - collection:
        name: Rocketblend
        platforms:
            - linux
            - windows
            - macos
        packages:
            - github.com/rocketblend/official-library/packages/rocketblend/0.1.0
```