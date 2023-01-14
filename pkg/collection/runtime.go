package collection

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rocketblend/rocketblend/pkg/core/runtime"
)

const (
	blenderExecutable = "blender"
	appContents       = "Blender.app/Contents/MacOS/"
)

var executableNames = map[runtime.Platform]string{
	runtime.Linux:     blenderExecutable,
	runtime.Windows:   blenderExecutable + ".exe",
	runtime.DarwinAmd: appContents + blenderExecutable,
	runtime.DarwinArm: appContents + blenderExecutable,
}

func getRuntimeExecutable(fileName string, platform runtime.Platform) (string, error) {
	executableName, ok := executableNames[platform]
	if !ok {
		return "", fmt.Errorf("executable not found for platform: %v", platform)
	}

	switch platform {
	case runtime.Windows, runtime.Linux:
		return filepath.Join(trimSuffix(fileName), executableName), nil
	default:
		return executableName, nil
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
	base := filepath.Base(fileName)
	return strings.TrimSuffix(base, filepath.Ext(base))
}
