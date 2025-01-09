package ssh

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/linyejoe2/gomni/utils"
)

// remote can be hostName or IP
func getRemote(remoteStr string) error {
	configPath := utils.GetFilePath()

	file, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read remote list: %w", err)
	}

	var remotes utils.Remotes
	if err := json.Unmarshal(file, &remotes); err != nil {
		return fmt.Errorf("failed to parse remote list: %w", err)
	}

	var remote *utils.Remote
	for _, e := range remotes.Remotes {
		if e.IP == remoteStr || e.Name == remoteStr {
			remote = &e
		}
	}

	if remote == nil {
		return fmt.Errorf("Can't find remote, please check your remote list by 'gomni ssh list'")
	}

	return connectRemote(remote)
}

func connectRemote(remote *utils.Remote) error {
	var cmd *exec.Cmd

	out, err := exec.Command("sshpass").CombinedOutput()
	if err != nil {
		if string(regexp.MustCompile(`No such file or directory`).Find(out)) != "" {
			return fmt.Errorf("Need to install sshpass first, run 'apt-get install sshpass'.")
		}
	}

	if remote.Auth.PrivateKey != "" {
		cmd = exec.Command("ssh", "-i", remote.Auth.PrivateKey, fmt.Sprintf("%s@%s", remote.Auth.Username, remote.IP))
	} else {
		cmd = exec.Command("sshpass", "-p", remote.Auth.Password, "ssh", fmt.Sprintf("%s@%s", remote.Auth.Username, remote.IP))
	}

	// bind std io to this terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
