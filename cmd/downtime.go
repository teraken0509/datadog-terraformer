package cmd

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/kterada0509/datadog-terraformer/internal/validations"
	middleware "github.com/kterada0509/datadog-terraformer/middleware/datadog"
)

// NewCmdDowntime ...
func NewCmdDowntime() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "downtime",
		Short: "Display 'datadog_downtime' resource configuration",
		Long:  `Display 'datadog_downtime' resource terraform configuration.`,
		Args:  validations.ValidationIntArg,
		RunE: func(cmd *cobra.Command, args []string) error {
			downtimeID, _ := strconv.Atoi(args[0])
			downtime, err := credential.GetDowntime(downtimeID)
			if err != nil {
				return err
			}
			if err = middleware.PrintDowntimeConfiguration(os.Stdout, downtime); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
