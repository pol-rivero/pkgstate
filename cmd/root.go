package cmd

import (
	"os"

	"github.com/pol-rivero/pkgstate/lib/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pkgstate",
	Short: "Declarative system state management.\nVersion: " + VERSION_STRING,
	Run: func(cmd *cobra.Command, args []string) {
		SetUpLogger(cmd)
		commands.Diff()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var basicCommandsGroup = &cobra.Group{
	ID:    "basicCommands",
	Title: "Basic commands:",
}

var otherCommandsGroup = &cobra.Group{
	ID:    "otherCommands",
	Title: "Other commands:",
}

func init() {
	rootCmd.Args = cobra.NoArgs
	rootCmd.AddGroup(basicCommandsGroup)
	rootCmd.AddGroup(otherCommandsGroup)
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Print additional information to stdout.")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress warnings and errors.")

	rootCmd.SetHelpCommandGroupID(otherCommandsGroup.ID)
	rootCmd.SetCompletionCommandGroupID(otherCommandsGroup.ID)
}
