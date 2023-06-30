package main

import (
	"fmt"
	"startup/apis"
	"startup/apps"
	"startup/cmd"
)

func main() {
	cmd.Execute()
}

func exec() {
	fmt.Println("Good Morning!")

	apps.OpenApplications()
	apis.VerifyAPIs()
}
