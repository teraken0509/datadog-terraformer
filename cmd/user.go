package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kterada0509/datadog-terraformer/internal/validations"
	middleware "github.com/kterada0509/datadog-terraformer/middleware/datadog"
)

// NewCmdUser ...
func NewCmdUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Display `user` configuration",
		Long:  "Display `user` terraform configuration.",
		Args:  validations.ValidationEmailArg,
		RunE: func(cmd *cobra.Command, args []string) error {
			user, err := credential.GetUser(args[0])
			if err != nil {
				return err
			}
			if err = middleware.PrintUserConfiguration(user); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
