/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/linyejoe2/gomni/cmd/ssh"
	"github.com/spf13/cobra"
)

var versionFlag bool
var version string = "dev"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gomni ssh",
	Short: "A command-line tool for routine work",
	Long: `gomni is a command-line tool for routine work,
it is very ease to learn and use!

feature:

 + gomni ssh: ssh client management tool.

for more please see https://github.com/linyejoe2/gomni
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			println("gomni", version)
			return
		}
		println(cmd.UsageString())
		return
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gomni.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	RootCmd.AddCommand(ssh.SshCmd)

	RootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "use to check the version of this tool.")

	// RootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	RootCmd.Root().CompletionOptions.HiddenDefaultCmd = true
}
