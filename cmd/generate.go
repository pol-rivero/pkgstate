package cmd

import (
	"github.com/pol-rivero/pkgstate/lib/commands"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	GroupID: basicCommandsGroup.ID,
	Use:     "generate",
	Short:   "Generates a configuration based on the current system state",
	Run: func(cmd *cobra.Command, args []string) {
		SetUpLogger(cmd)
		commands.Generate()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Args = cobra.NoArgs
}
