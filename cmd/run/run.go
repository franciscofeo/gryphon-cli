/*
Copyright © 2023 Francisco Angelo <franciscoangelo.dev@gmail.com>
*/

package run

import (
	"github.com/spf13/cobra"
)

// RunCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run is a palette that execute all modules of the application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
}
