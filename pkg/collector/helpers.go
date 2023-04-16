package collector

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rocketblend/rocketblend/pkg/rocketblend/runtime"
	"github.com/rocketblend/rocketblend/pkg/semver"
)

const (
	WindowsPlatformRegex string = "(blender-.+win.+64.+zip)"
	LinuxPlatformRegex   string = "(blender-.+lin.+64.+tar)"
	MacPlatformRegex     string = "(blender-.+(macos|darwin).+dmg)"
	ArmMacPlatformRegex  string = "(arm64)"
	VersionNumberRegex   string = "[0-9]+([.][0-9]+)"
)

func findVerisonNumberStr(input string) string {
	// Gets the full version number from the input string.
	r, _ := regexp.Compile(fmt.Sprintf("(%s+)", VersionNumberRegex))
	versionStr := r.FindString(strings.ToLower(input))
	return versionStr
}

func parseVersionNumber(input string) (*semver.Version, error) {
	// Gets the full version number from the input string.
	versionStr := trimOrPad(findVerisonNumberStr(input))
	return semver.Parse(versionStr)
}

func trimOrPad(s string) string {
	// Split the string into parts.
	parts := strings.Split(s, ".")

	// Pad or trim the parts as needed.
	major := parts[0]
	minor := "0"
	patch := "0"
	if len(parts) > 1 {
		minor = parts[1]
	}
	if len(parts) > 2 {
		patch = parts[2]
	}

	// Return the padded or trimmed string.
	return fmt.Sprintf("%s.%s.%s", major, minor, patch)
}

func parseMajorMinorVersionNumber(input string) (float32, error) {
	// Gets just the major and minor version number from the input string.
	r, _ := regexp.Compile(VersionNumberRegex)
	versionStr := r.FindString(strings.ToLower(input))
	value, err := strconv.ParseFloat(versionStr, 32)
	if err != nil {
		return 0, err
	}

	return float32(value), nil
}

func parsePlatform(name string) runtime.Platform {
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

func censorText(text string, char string, limit int) string {
	// Check if the limit is greater than the length of the input text
	if limit > len(text) {
		limit = len(text)
	}

	// Create a new string with the specified number of characters
	return strings.Repeat(char, limit)
}
