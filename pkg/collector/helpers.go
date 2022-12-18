package collector

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rocketblend/rocketblend/pkg/core/runtime"
)

const (
	WindowsPlatformRegex string = "(blender-.+win.+64.+zip)"
	LinuxPlatformRegex   string = "(blender-.+lin.+64.+tar)"
	MacPlatformRegex     string = "(blender-.+(macos|darwin).+dmg)"
	ArmMacPlatformRegex  string = "(arm64)"
	VersionNumberRegex   string = "[0-9]+([.][0-9]+)"
)

var Platforms = [4]string{"windows", "linux", "macos arm", "macos intel"}

func FindVerisonNumberStr(input string) string {
	// Gets the full version number from the input string.
	r, _ := regexp.Compile(fmt.Sprintf("(%s+)", VersionNumberRegex))
	versionStr := r.FindString(strings.ToLower(input))
	return versionStr
}

func ParseMajorMinorVersionNumber(input string) (float32, error) {
	// Gets just the major and minor version number from the input string.
	r, _ := regexp.Compile(VersionNumberRegex)
	versionStr := r.FindString(strings.ToLower(input))
	value, err := strconv.ParseFloat(versionStr, 32)
	if err != nil {
		return 0, err
	}

	return float32(value), nil
}

func ParsePlatform(name string) runtime.Platform {
	name = strings.ToLower(name)
	match, _ := regexp.MatchString(WindowsPlatformRegex, name)
	if match {
		return runtime.Windows
	}

	match, _ = regexp.MatchString(LinuxPlatformRegex, name)
	if match {
		return runtime.Linux
	}

	match, _ = regexp.MatchString(MacPlatformRegex, name)
	if match {
		// regexp packages doesn't support look-ahead/look-behind
		IsArm, _ := regexp.MatchString(ArmMacPlatformRegex, name)
		if IsArm {
			return runtime.DarwinArm
		}

		return runtime.DarwinAmd
	}

	return runtime.Undefined
}

func GenerateHash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
