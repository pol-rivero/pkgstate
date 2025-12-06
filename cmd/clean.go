package cmd

import (
	"github.com/pol-rivero/pkgstate/lib/commands"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	GroupID: basicCommandsGroup.ID,
	Use:     "clean",
	Short:   "Cleans up all orphaned and optional packages. Warning: This might break things.",
	Run: func(cmd *cobra.Command, args []string) {
		SetUpLogger(cmd)
		noConfirm, err := cmd.Flags().GetBool("yes")
		if err != nil {
			panic(err)
		}
		commands.Clean(noConfirm)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Args = cobra.NoArgs
	cleanCmd.Flags().BoolP("yes", "y", false, "Apply changes without confirmation.")
}
