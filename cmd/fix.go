package cmd

import (
	"github.com/pol-rivero/pkgstate/lib/commands"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	GroupID: basicCommandsGroup.ID,
	Use:     "fix",
	Short:   "Makes the required changes for the system to match the desired state defined in the config files",
	Run: func(cmd *cobra.Command, args []string) {
		SetUpLogger(cmd)
		noConfirm, err := cmd.Flags().GetBool("yes")
		if err != nil {
			panic(err)
		}
		noRemove, err := cmd.Flags().GetBool("no-remove")
		if err != nil {
			panic(err)
		}
		commands.Fix(noConfirm, noRemove)
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)

	fixCmd.Args = cobra.NoArgs
	fixCmd.Flags().BoolP("yes", "y", false, "Apply changes without confirmation.")
	fixCmd.Flags().Bool("no-remove", false, "Do not uninstall any packages, even when --yes is passed.")
}
