package collection

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/rocketblend/rocketblend/pkg/core/runtime"
)

func getRuntimeExecutable(fileName string, platform runtime.Platform) string {
	switch platform {
	case runtime.Windows:
		return path.Join(trimSuffix(fileName), "blender.exe")
	case runtime.DarwinAmd, runtime.DarwinArm:
		return "Blender.app/Contents/MacOS/blender"
	default:
		return "blender"
	}
}

func contains(s []runtime.Platform, e runtime.Platform) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func trimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
