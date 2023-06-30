/*
Copyright © 2023 Francisco Angelo <franciscoangelo.dev@gmail.com>
*/

package cmd

import (
	"os"
	"startup/cmd/api"
	"startup/cmd/app"
	"startup/cmd/run"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gryphon",
	Short: "A simple CLI to bring some information and start main apps",
	Long: `⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣠⠀⠂⠠⢀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠈⠐⢄⠀⠑⡩⢂⠀⠀⠀⠀⠀⠀Gryphon Startup is an application to bring information about the main APIs used in the web development world: GitHub, Slack, Atlassian and more.⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⡀⠀⢁⠀⢀⡆⠃⠀⠀⠀⠀⠀Furthermore, it starts essentials apps with a single command to improve your time.⠀⠀⠀⠀⠀⠀⠀
⠀⢀⠀⠔⠂⠂⠅⠀⠘⠀⢤⠀⣄⠆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠈⠈⠐⠀⠀⢨⠠⢁⠜⠁⢔⠌⠆⠀⠀⠀⠀ Author:	Francisco Angelo⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢀⠊⠀⠀⠀⠀⠀⠠⠒⡡⠊⠀⠀⠀⠀⠀ LinkedIn:	linkedin.com/in/francisco-angelo⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠸⠀⠀⠀⡀⠀⠀⠠⡈⠀⢐⠔⢐⠠⡀⠀⠀Email:	franciscoangelo.dev@gmail.com⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠑⠀⣊⠀⠀⠀⠀⠈⠢⠘⠇⠈⠐⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢀⠀⡱⠈⢂⠀⠀⠀⠀⡠⠰⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠰⠭⠤⠌⠀⠰⠬⠤⠥⠤⠤⠄⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.startup.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommandPalettes()
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(app.AppCmd)
	rootCmd.AddCommand(run.RunCmd)
	rootCmd.AddCommand(api.ApiCmd)
}
