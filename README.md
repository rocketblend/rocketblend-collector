# RocketBlend Collector

[![Github tag](https://badgen.net/github/tag/rocketblend/rocketblend-collector)](https://github.com/rocketblend/rocketblend-collector/tags)
[![Go Doc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/rocketblend/rocketblend-collector)
[![Go Report Card](https://goreportcard.com/badge/github.com/rocketblend/rocketblend-collector)](https://goreportcard.com/report/github.com/rocketblend/rocketblend-collector)
[![GitHub](https://img.shields.io/github/license/rocketblend/rocketblend-collector)](https://github.com/rocketblend/rocketblend-collector/blob/master/LICENSE)

CLI tool for collecting [Blender](https://www.blender.org/) build information and generating [Libraries](https://github.com/rocketblend/official-library) for use with use with [RocketBlend](https://github.com/rocketblend/rocketblend)

## Example config

```yaml
library: github.com/rocketblend/official-library
collector:
  proxy: http://user:pass@proxy.com
  agent: random
  parallelism: 2
  delay: 15s
collections:
  - name: blender
    platforms:
      - windows
      - linux
      - macos/intel
      - macos/apple
  - name: rocketblend
    platforms:
      - windows
    packages:
      - github.com/rocketblend/official-library/packages/rocketblend/0.1.0
```

You can also use a environment variable to set the proxy url.

```bash
export COLLECTOR_PROXY="http://user:pass@proxy.com"
```

## Usage

```bash
collector pull
```