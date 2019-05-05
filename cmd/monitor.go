package cmd

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/kterada0509/datadog-terraformer/internal/validations"
	middleware "github.com/kterada0509/datadog-terraformer/middleware/datadog"
)

// NewCmdMonitor ...
func NewCmdMonitor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Display 'datadog_monitor' resource configuration",
		Long:  `Display 'datadog_monitor' resource terraform configuration.`,
		Args:  validations.ValidationIntArg,
		RunE: func(cmd *cobra.Command, args []string) error {
			monitorID, _ := strconv.Atoi(args[0])
			monitor, err := credential.GetMonitor(monitorID)
			if err != nil {
				return err
			}
			if err = middleware.PrintMonitorConfiguration(os.Stdout, monitor); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
