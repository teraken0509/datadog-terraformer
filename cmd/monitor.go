package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	middleware "github.com/kterada0509/datadog-terraformer/middleware/datadog"
)

// NewCmdMonitor ...
func NewCmdMonitor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Display `monitor` configuration",
		Long:  "Display `monitor` terraform configuration.",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("requires one arg")
			}
			if _, err := strconv.Atoi(args[0]); err != nil {
				return fmt.Errorf("Invalid argument expect int but got %s", args[0])
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			monitorID, _ := strconv.Atoi(args[0])
			monitor, err := credential.GetMonitor(monitorID)
			if err != nil {
				return err
			}
			if err = middleware.PrintMonitorConfiguration(monitor); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
