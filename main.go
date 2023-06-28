package main

import (
	"fmt"
	"startup/api"
	"startup/apps"
)

func main() {
	fmt.Println("Good Morning!")

	apps.OpenApplications()
	api.VerifyAPIs()
}
