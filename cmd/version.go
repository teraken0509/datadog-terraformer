package cmd

import (
	"github.com/kterada0509/datadog-terraformer/internal"
	"github.com/spf13/cobra"
)

// NewCmdVersion ...
func NewCmdVersion() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "print command line tool version",
		Long:  "print command line tool version",
		Run: func(cmd *cobra.Command, args []string) {
			internal.PrintVersion()
		},
	}

	return cmd
}
