package validations

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// ValidationIntArg ...
func ValidationIntArg(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Unexpected number of arguments. expected: 1, got: %d", len(args))
	}
	if _, err := strconv.Atoi(args[0]); err != nil {
		return fmt.Errorf("Invalid argument expect int but got %s", args[0])
	}

	return nil
}

// ValidationEmailArg ...
func ValidationEmailArg(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Unexpected number of arguments. expected: 1, got: %d", len(args))
	}

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(args[0]) {
		return fmt.Errorf("Invalid argument expect email format but got %s", args[0])
	}

	return nil
}
