/*
Copyright © 2023 Francisco Angelo <franciscoangelo.dev@gmail.com>
*/

package cmd

import (
	"gryphon/cmd/api"
	"gryphon/cmd/app"
	"gryphon/cmd/run"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gryphon",
	Short: "A simple CLI to bring some information and start main apps",
	Long: `⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⠀⠀⣀⡀⠀⠀⢀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣼⣦⠀⠀⢀⣾⣿⠀⢠⣽⣿⣿⣿⣿⣟⣿⣤⣀⡀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣿⣿⣷⠀⣾⣿⣿⡄⠀⢨⣿⣿⣿⣿⣿⣿⣟⡛⠛⠃⠀⠀⠀⠀⠀⠀Gryphon CLI is an application to bring information about the main APIs used in the web development world: GitHub, Slack, Atlassian and more.⠀
⠀⠀⠀⠀⣿⣿⡇⢠⣿⣿⣿⣿⡀⢸⣿⣿⣿⣿⣿⠛⠛⠛⠛⠀⠀⠀⠀⠀⠀⠀Furthermore, it starts essentials apps with a single command to improve your time.⠀
⠀⠀⠀⠀⣿⣿⡇⢸⣿⣿⣿⣿⣷⡜⣿⣿⣿⣿⣿⣷⣦⣀⠀⠀⠀⢠⣶⠦⠄⠀ 
⠀⢀⡀⠀⠹⣿⡇⠸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣀⣤⣾⣅⣀⠀⠀ Author:	 Francisco Angelo⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠈⣿⡆⠀⠈⠻⠀⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠛⠉⠉⠁⠀⠀ LinkedIn: linkedin.com/in/francisco-angelo⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠉⠙⠦⣄⡀⠀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣏⡀⠀⠀⣀⠀⠀⠀ Email:	 franciscoangelo.dev@gmail.com⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠈⣻⡆⠀⠀⠈⠙⣛⣿⣿⣿⣿⣿⣿⣿⣿⡿⢿⣿⣷⣿⠛⠓⠀⠀
⠀⠀⠀⠀⢀⣾⠟⠁⠀⢀⣰⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠀⠀⠀⠈⠙⠿⠀⠀⠀
⠀⠀⠀⠀⣿⣇⣀⣤⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠉⠛⠛⠛⠉⢁⣾⣿⣿⣿⡟⠉⢻⣿⣿⣇⣀⠀⢀⣀⣀⡀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣤⣶⣿⣿⣿⣿⠟⠀⠀⠀⢻⣿⠿⠿⣷⣿⣏⠉⠙⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢀⣼⣿⣿⣭⡉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣹⠇⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠈⠙⠛⠉⠀⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommandPalettes()
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(app.AppCmd)
	rootCmd.AddCommand(run.RunCmd)
	rootCmd.AddCommand(api.ApiCmd)
}
