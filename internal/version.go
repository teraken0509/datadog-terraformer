package internal

import (
	"fmt"
)

// version The main version number
var version string

// PrintVersion  returns the complete version string, including prerelease
func PrintVersion() string {
	return fmt.Sprintf("%s %s", Package, version)
}
