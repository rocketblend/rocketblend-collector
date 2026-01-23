package collector

import (
	"testing"

	"github.com/rocketblend/rocketblend/pkg/runtime"
)

func TestParsePlatform(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     runtime.Platform
	}{
		// Windows x64 - should match
		{
			name:     "Windows x64 v5.0.0 zip",
			filename: "blender-5.0.0-windows-x64.zip",
			want:     runtime.Windows,
		},
		{
			name:     "Windows x64 v4.2.13 zip",
			filename: "blender-4.2.13-windows-x64.zip",
			want:     runtime.Windows,
		},
		{
			name:     "Windows x64 v5.0.1 zip",
			filename: "blender-5.0.1-windows-x64.zip",
			want:     runtime.Windows,
		},
		// Old Windows format (pre-v3) - should match
		{
			name:     "Windows v2.92.0 old format",
			filename: "blender-2.92.0-windows64.zip",
			want:     runtime.Windows,
		},
		{
			name:     "Windows v2.83.0 old format",
			filename: "blender-2.83.0-windows64.zip",
			want:     runtime.Windows,
		},
		// Windows ARM64 - should NOT match (returns Undefined)
		{
			name:     "Windows ARM64 v5.0.0 zip",
			filename: "blender-5.0.0-windows-arm64.zip",
			want:     runtime.Undefined,
		},
		{
			name:     "Windows ARM64 v5.0.1 zip",
			filename: "blender-5.0.1-windows-arm64.zip",
			want:     runtime.Undefined,
		},
		// Linux x64 - should match
		{
			name:     "Linux x64 v5.0.0 tar.xz",
			filename: "blender-5.0.0-linux-x64.tar.xz",
			want:     runtime.Linux,
		},
		{
			name:     "Linux x64 v4.2.13 tar.xz",
			filename: "blender-4.2.13-linux-x64.tar.xz",
			want:     runtime.Linux,
		},
		// macOS ARM64 - should match
		{
			name:     "macOS ARM64 v5.0.0 dmg",
			filename: "blender-5.0.0-macos-arm64.dmg",
			want:     runtime.DarwinArm,
		},
		{
			name:     "macOS ARM64 v4.2.13 dmg",
			filename: "blender-4.2.13-macos-arm64.dmg",
			want:     runtime.DarwinArm,
		},
		// macOS x64 (Intel) - should match (no arm64 in name)
		{
			name:     "macOS x64 v5.0.0 dmg (implicit)",
			filename: "blender-5.0.0-macos-x64.dmg",
			want:     runtime.DarwinAmd,
		},
		{
			name:     "macOS x64 v4.2.13 dmg (implicit)",
			filename: "blender-4.2.13-macos-x64.dmg",
			want:     runtime.DarwinAmd,
		},
		// Other file types that shouldn't match
		{
			name:     "Windows MSI - should not match (only zip)",
			filename: "blender-5.0.0-windows-x64.msi",
			want:     runtime.Undefined,
		},
		{
			name:     "Windows MSIX - should not match (only zip)",
			filename: "blender-5.0.0-windows-x64.msix",
			want:     runtime.Undefined,
		},
		{
			name:     "MD5 checksum file - should not match",
			filename: "blender-5.0.0.md5",
			want:     runtime.Undefined,
		},
		{
			name:     "SHA256 checksum file - should not match",
			filename: "blender-5.0.0.sha256",
			want:     runtime.Undefined,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parsePlatform(tt.filename)
			if got != tt.want {
				t.Errorf("parsePlatform(%q) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}

func TestParseVersionNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "Version 5.0.0",
			input:   "blender-5.0.0-windows-x64.zip",
			want:    "5.0.0",
			wantErr: false,
		},
		{
			name:    "Version 4.2.13",
			input:   "blender-4.2.13-linux-x64.tar.xz",
			want:    "4.2.13",
			wantErr: false,
		},
		{
			name:    "Version 2.83",
			input:   "blender-2.83-windows-x64.zip",
			want:    "2.83.0",
			wantErr: false,
		},
		{
			name:    "Version 3.6.7",
			input:   "blender-3.6.7-macos-arm64.dmg",
			want:    "3.6.7",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseVersionNumber(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseVersionNumber(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("parseVersionNumber(%q) = %v, want %v", tt.input, got.String(), tt.want)
			}
		})
	}
}
