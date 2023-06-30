/*
Copyright © 2023 Francisco Angelo <franciscoangelo.dev@gmail.com>
*/

package api

import (
	"github.com/spf13/cobra"
)

// ApiCmd represents the apis command
var ApiCmd = &cobra.Command{
	Use:   "apis",
	Short: "Api is a palette that contains API check commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
