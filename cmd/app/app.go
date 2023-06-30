/*
Copyright © 2023 Francisco Angelo <franciscoangelo.dev@gmail.com>
*/

package app

import (
	"github.com/spf13/cobra"
	"startup/apps"
)

const (
	listFlagName      = "list"
	listFlagShorthand = "l"
)

// AppCmd represents the app command
var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "App is a palette that contains application runner commands",
	Long:  ``,
	Run:   main,
}

func main(cmd *cobra.Command, args []string) {
	showList, _ := cmd.Flags().GetBool(listFlagName)
	if showList {
		apps.ListApplications()
	} else {
		_ = cmd.Help()
	}
}

func init() {
	AppCmd.Flags().BoolP(listFlagName, listFlagShorthand, false, "List all available apps")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
