package ssh

import (
	"fmt"
	"os"
	"strings"

	"github.com/linyejoe2/gomni/utils"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
	hostName string
	keyFile  string
)

func validateKeyFile(keyFile string) error {
	info, err := os.Stat(keyFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	if info.IsDir() {
		return fmt.Errorf("path is a directory, not a file")
	}

	content, err := os.ReadFile(keyFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	if !isValidPrivateKey(content) {
		return fmt.Errorf("file is not a valid SSH private key")
	}

	return nil
}

func isValidPrivateKey(content []byte) bool {
	keyString := string(content)
	return strings.HasPrefix(keyString, "-----BEGIN OPENSSH PRIVATE KEY-----") ||
		strings.HasPrefix(keyString, "-----BEGIN RSA PRIVATE KEY-----")
}

var addCmd = &cobra.Command{
	Use: `add <ip> -n <hostname> -u <username> -p <password> or -i <certificate file>

💡example:
 + with password: gomni ssh add 127.0.0.1 -n my-pc -u user -p 1234
 + with certificate file: gomni ssh add 127.0.0.1 -n my-pc -u user -i ~/.ssh/id_rsa

❗note:
 + do not use reserved word like "add, delete, list" as hostname, because those word will cause panic error!
 + if you use certificate, you must add public key to server first, 
   or ssh will eithre fall back to use password (if password log-in is supported), 
   or print fail message and close.
 + certificate file is private key, not public key!
	`,
	Short: "Add a new SSH remote.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := args[0]
		if hostName == "" {
			fmt.Println("Invalid input, must provide hostname with -n flag.")
			return
		}

		if username == "" {
			fmt.Println("Invalid input, must provide username with -u flag.")
			return
		}

		if password == "" && keyFile == "" {
			fmt.Println("Invalid input, must provide either password with -p flag or certificate with -i flag.")
			return
		}

		if password != "" {
			fmt.Printf("Adding remote %s with password auth\n", ip)
		} else {
			fmt.Printf("Adding remote %s with certificate\n", ip)
		}

		err := addRemote(ip, hostName, username, password, keyFile)
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
	addCmd.Flags().StringVarP(&hostName, "name", "n", "", "Remote name")
}
