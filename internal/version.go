package internal

import (
	"fmt"
)

// version The main version number
var version, date string

// PrintVersion  returns the complete version string, including prerelease
func PrintVersion() string {
	return fmt.Sprintf("%s %s, Release Date: %s", Package, version, date)
}
