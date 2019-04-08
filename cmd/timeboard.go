package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/kterada0509/datadog-terraformer/internal/validations"
	middleware "github.com/kterada0509/datadog-terraformer/middleware/datadog"
)

// NewCmdTimeboard ...
func NewCmdTimeboard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "timeboard",
		Short: "Display `timeboard` configuration",
		Long:  "Display `timeboard` terraform configuration.",
		Args:  validations.ValidationIntArg,
		RunE: func(cmd *cobra.Command, args []string) error {
			boardID, _ := strconv.Atoi(args[0])
			board, err := credential.GetTimeboard(boardID)
			if err != nil {
				return err
			}
			if err = middleware.PrintTimeBoardConfiguration(board); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
