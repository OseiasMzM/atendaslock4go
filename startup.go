package main

import (
	"atendaslock4go/lock4"
	"atendaslock4go/logs"
	"atendaslock4go/router"
)

func InitializeServer() {
	logs.AddText("Server Started...")
	logs.AddText("Version:" + lock4.LOCK4_VERSION)

	router.Initialize()
}
