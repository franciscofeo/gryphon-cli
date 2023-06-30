package main

import (
	"fmt"
	"startup/api"
	"startup/apps"
	"startup/cmd"
)

func main() {
	cmd.Execute()
}

func exec() {
	fmt.Println("Good Morning!")

	apps.OpenApplications()
	api.VerifyAPIs()
}
