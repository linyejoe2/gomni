package ssh

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	"github.com/linyejoe2/gomni/utils"
	"github.com/spf13/cobra"
)

func listRemote() ([]utils.Remote, error) {
	configPath := utils.GetFilePath()

	file, err := os.ReadFile(configPath)
	if err != nil {
		return *new([]utils.Remote), fmt.Errorf("failed to read remote list: %w", err)
	}

	var remotes utils.Remotes
	if err := json.Unmarshal(file, &remotes); err != nil {
		return *new([]utils.Remote), fmt.Errorf("failed to parse remote list: %w", err)
	}

	return remotes.Remotes, nil
}

func createPrintList(remotes []utils.Remote) (list []string) {
	// red := color.New(color.FgRed).SprintFunc()
	// green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	for _, remote := range remotes {
		list = append(list, remote.Name+"\t "+remote.IP+"\t "+remote.Auth.Username+"\t "+yellow("panding"))
	}

	return
}

func fPrintHelper(w *tabwriter.Writer, list []string) {
	fmt.Fprintln(w, "NAME\t IP\t USERNAME\t STATUS\t")

	for _, s := range list {
		fmt.Fprintln(w, s)
	}
	w.Flush()
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all avaliable remotes",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		remotes, err := listRemote()
		if err != nil {
			fmt.Println("Can't list remotes: ", err)
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		printList := createPrintList(remotes)

		fPrintHelper(w, printList)

		time.Sleep(1 * time.Second)

		for i, remote := range remotes {
			var status string
			if utils.CheckSSHRemoteAlive(remote.IP) {
				status = green("online ")
			} else {
				status = red("offline")
			}

			printList = append(append(printList[:i], remote.Name+"\t "+remote.IP+"\t "+remote.Auth.Username+"\t "+status), printList[i+1:]...)

			utils.ClearStdOutPreLine(len(remotes) + 1)

			fPrintHelper(w, printList)
		}

	},
}

func init() {
	SshCmd.AddCommand(listCmd)
}
