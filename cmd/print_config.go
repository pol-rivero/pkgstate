package cmd

import (
	printconfig "github.com/pol-rivero/pkgstate/lib/commands/print_config"
	"github.com/spf13/cobra"
)

var printConfigCmd = &cobra.Command{
	GroupID: otherCommandsGroup.ID,
	Use:     "print-config",
	Short:   "Prints the effective configuration resulting from merging all config files. This can be useful for debugging.",
	Run: func(cmd *cobra.Command, args []string) {
		SetUpLogger(cmd)
		printconfig.PrintConfig()
	},
}

func init() {
	rootCmd.AddCommand(printConfigCmd)
	printConfigCmd.Args = cobra.NoArgs
}
