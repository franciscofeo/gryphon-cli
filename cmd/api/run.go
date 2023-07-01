/*
Copyright © 2023 Francisco Angelo <franciscoangelo.dev@gmail.com>
*/

package api

import (
	"github.com/spf13/cobra"
	"gryphon/apis"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "This command verify all the APIs",
	Long:  ``,
	Run:   verifyAll,
}

func verifyAll(cmd *cobra.Command, args []string) {
	apis.VerifyAPIs()
}

func init() {
	ApiCmd.AddCommand(runCmd)
}
