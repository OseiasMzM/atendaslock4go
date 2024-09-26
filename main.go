package main

import (
	"atendaslock4go/database"
	"atendaslock4go/paths"
	"fmt"
)

func main() {

	fmt.Println("Started...")

	paths.InstanceDataBase()
	database.Config()

	switch ParseParams() {
	case "/install":
		InstallService()
	case "/uninstall":
		UninstallService()
	default:
		RunService()
	}

}
