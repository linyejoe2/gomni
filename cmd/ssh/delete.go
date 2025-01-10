package ssh

import (
	"fmt"
	"os"

	"github.com/linyejoe2/gomni/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: `delete <remote name | ip>

	💡example:
	 + let assume that you have remote list 
	NAME             | IP            | USERNAME	| STATUS   |
	my-pc            | 127.0.0.1	 | user		| online   |
	 + you can delete my-pc by:
	  + gomni ssh delete my-pc
	  + gomni ssh delete 127.0.0.1
	`,
	Short: "Delete specify remote.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := utils.GetFilePath()

		// 讀取現有 remotes
		remotes, err := utils.LoadRemotes(filePath)
		if err != nil && !os.IsNotExist(err) {
			fmt.Println("Can't get remotes: ", err.Error())
		}

		for i, remote := range remotes.Remotes {
			if remote.IP == args[0] || remote.Name == args[0] {
				remotes.Remotes = append(remotes.Remotes[:i], remotes.Remotes[i+1:]...)
				utils.SaveRemotes(utils.GetFilePath(), remotes)
				fmt.Println("Successfully delete remote", remote.Name)
				return
			}
		}

		fmt.Println("Can't find remote '", args[0], "', did you enter the right remote name or ip?")
	},
}

func init() {
	SshCmd.AddCommand(deleteCmd)
}
