package collection

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/rocketblend/rocketblend/pkg/driver/runtime"
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

	executablePath := ""
	switch platform {
	case runtime.Windows:
		executablePath = trimExt(fileName)
	case runtime.Linux:
		executablePath = trimExt(trimExt(fileName))
	}

	return filepath.Join(executablePath, executableName), nil
}

func contains(s []runtime.Platform, e runtime.Platform) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func trimExt(fileName string) string {
	if fileName == "" || !strings.Contains(fileName, ".") {
		return fileName
	}

	return strings.TrimSuffix(path.Base(fileName), path.Ext(fileName))
}
