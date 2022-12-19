package collection

import (
	"fmt"

	"github.com/rocketblend/rocketblend/pkg/core/runtime"
)

const BlenderExecutable = "blender"

func getRuntimeExecutable(platform runtime.Platform) string {
	switch platform {
	case runtime.Windows:
		return fmt.Sprintf("%s.exe", BlenderExecutable)
	case runtime.DarwinAmd, runtime.DarwinArm:
		return fmt.Sprintf("%s.dmg", BlenderExecutable)
	default:
		return BlenderExecutable
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
