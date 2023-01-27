### [Discussions](https://github.com/rocketblend/rocketblend-collector/discussions) │ [Documentation](https://docs.rocketblend.io/v/collector) │ [Latest Release](https://github.com/rocketblend/rocketblend-collector/releases/latest)

# RocketBlend Collector

[![Github tag](https://badgen.net/github/tag/rocketblend/rocketblend-collector)](https://github.com/rocketblend/rocketblend-collector/tags)
[![Go Doc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/rocketblend/rocketblend-collector)
[![Go Report Card](https://goreportcard.com/badge/github.com/rocketblend/rocketblend-collector)](https://goreportcard.com/report/github.com/rocketblend/rocketblend-collector)
[![GitHub](https://img.shields.io/github/license/rocketblend/rocketblend-collector)](https://github.com/rocketblend/rocketblend-collector/blob/master/LICENSE)


![Hero image of Collector CLI](docs/assets/collector-about.svg)

> Collector is a command-line tool for generating package configurations quickly and easily for use with [RocketBlend](https://github.com/rocketblend/rocketblend).

## Example config

```yaml
# collector.yaml
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
    addons:
      - github.com/rocketblend/official-library/addons/rocketblend/0.1.0
```

You can also use an environment variable to set the proxy url.

```bash
export COLLECTOR_PROXY="http://user:pass@proxy.com"
```

## Usage

```bash
collector pull
```

## See Also

- [RocketBlend](https://github.com/rocketblend/rocketblend) - Main project.
- [RocketBlend Launcher](https://github.com/rocketblend/rocketblend-launcher) - Replacement launcher for Blender that utilises RocketBlend.
- [RocketBlend Companion](https://github.com/rocketblend/rocketblend-companion) - Blender addon to aid with working with RocketBlend. **NOTE: WIP**
- [Official Library](https://github.com/rocketblend/official-library) - Collection of builds and addons.
