/*
Copyright © 2023 Francisco Angelo <franciscoangelo.dev@gmail.com>
*/

package api

import (
	"github.com/spf13/cobra"
	"gryphon/apis"
)

const (
	listFlagName  = "list"
	listShorthand = "l"
)

// ApiCmd represents the apis command
var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Api is a palette that contains API check commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		showAPIs, _ := cmd.Flags().GetBool(listFlagName)

		if showAPIs {
			apis.ListAvailableAPIs()
		} else {
			_ = cmd.Help()
		}
	},
}

func init() {
	ApiCmd.Flags().BoolP(listFlagName, listShorthand, false, "List all available APIs to verify")
}
