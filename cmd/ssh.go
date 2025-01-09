package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage SSH remotes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SSH command")
	},
}

func init() { rootCmd.AddCommand(sshCmd) }
