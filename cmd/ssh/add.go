package ssh

import (
	"fmt"
	"os"

	"github.com/linyejoe2/gomni/utils"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
	name     string
	keyFile  string
)

var addCmd = &cobra.Command{
	Use:   "add <ip> -u <username> -p <password> or -i <certificate file>",
	Short: "Add a new SSH remote",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := args[0]
		if username != "" && password != "" {
			fmt.Printf("Adding remote %s with password auth\n", ip)
		} else if keyFile != "" {
			fmt.Printf("Adding remote %s with certificate\n", ip)
		} else {
			fmt.Println("Invalid input, must provide either username/password with or certificate")
			return
		}
		err := addRemote(ip, name, username, password, keyFile)
		if err != nil {
			panic(err)
		}
	},
}

func addRemote(ip, name, username, password, keyFile string) error {
	filePath := utils.GetFilePath()

	// 讀取現有 remotes
	remotes, err := utils.LoadRemotes(filePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	for _, e := range remotes.Remotes {
		if e.IP == ip {
			fmt.Printf("Already have this remote, you need to 'gomni ssh remove %s' first.\n", ip)
			return nil
		}
	}

	// 新增 remote
	newRemote := utils.Remote{
		IP:   ip,
		Name: name,
	}
	newRemote.Auth.Username = username
	newRemote.Auth.Password = password
	newRemote.Auth.PrivateKey = keyFile

	remotes.Remotes = append(remotes.Remotes, newRemote)

	// 儲存回文件
	return utils.SaveRemotes(filePath, remotes)
}

func init() {
	SshCmd.AddCommand(addCmd)

	// 設定 flags
	addCmd.Flags().StringVarP(&username, "username", "u", "", "Username for SSH")
	addCmd.Flags().StringVarP(&password, "password", "p", "", "Password for SSH")
	addCmd.Flags().StringVarP(&keyFile, "identify", "i", "", "Private key file for SSH")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Remote name")
}
