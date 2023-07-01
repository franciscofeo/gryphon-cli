/*
Copyright © 2023 Francisco Angelo <franciscoangelo.dev@gmail.com>
*/

package app

import (
	"github.com/spf13/cobra"
	"startup/apps"
)

const (
	nameFlag      = "name"
	nameShortHand = "n"
)

var appName string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "This command opens all applications and URLs or you can specify the name to open a single app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if appName == "" {
			apps.OpenApplications()
			return
		}
		err := apps.OpenSingleApplication(appName)
		if err != nil {
			cmd.Println(err)
			_ = cmd.Help()
		}
	},
}

func init() {
	AppCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(
		&appName,
		nameFlag,
		nameShortHand,
		"",
		"Insert the name to open a single application",
	)
}
