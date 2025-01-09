package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SshCmd = &cobra.Command{
	Use:   "ssh [host name | ip]",
	Short: "Connect & Manage SSH remotes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		remoteName := args[0]

		// 呼叫 ConnectRemote 來進行連接
		if err := getRemote(remoteName); err != nil {
			return fmt.Errorf("failed to connect: %w", err)
		}

		return nil
	},
}
