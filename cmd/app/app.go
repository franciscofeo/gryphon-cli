/*
Copyright © 2023 Francisco Angelo <franciscoangelo.dev@gmail.com>
*/

package app

import (
	"github.com/spf13/cobra"
)

// AppCmd represents the app command
var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "App is a palette that contains application runner commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	AppCmd.Flags().BoolP("list", "l", false, "List all available apps")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
