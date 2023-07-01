/*
Copyright © 2023 Francisco Angelo <franciscoangelo.dev@gmail.com>
*/

package run

import (
	"github.com/spf13/cobra"
	"gryphon/apis"
	"gryphon/apps"
	"gryphon/utils"
	"strconv"
	"time"
)

const (
	sleepDurationFlagName  = "sleep"
	sleepDurationShorthand = "s"
	defaultDuration        = 400
)

var sleepDuration string

// RunCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run is a command that execute all modules of the application",
	Long:  ``,
	Run:   runAll,
}

func runAll(cmd *cobra.Command, args []string) {
	apis.VerifyAPIs()
	sleep(sleepDuration)
	apps.OpenApplications()
}

func sleep(duration string) {
	if utils.IsStringBlank(duration) {
		time.Sleep(defaultDuration * time.Millisecond)
	} else {
		value, _ := strconv.Atoi(duration)
		if value <= 600 {
			time.Sleep(time.Duration(value) * time.Millisecond)
		} else {
			time.Sleep(defaultDuration * time.Millisecond)
		}
	}
}

func init() {
	RunCmd.Flags().StringVarP(
		&sleepDuration,
		sleepDurationFlagName,
		sleepDurationShorthand,
		"",
		"Defines a delay duration between the execution of different modules",
	)
}
