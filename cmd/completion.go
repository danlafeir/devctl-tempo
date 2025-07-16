package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var completionCmd = &cobra.Command{
	Use:    "completion [bash|zsh|fish|powershell]",
	Short:  "Generate shell completion scripts",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Println("Please specify a shell: bash, zsh, fish, or powershell")
			os.Exit(1)
		}
		switch args[0] {
		case "bash":
			rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			cmd.Println("Unsupported shell type.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
