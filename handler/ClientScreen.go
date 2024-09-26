package handler

import (
	"atendaslock4go/paths"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func ClientScreen(c *gin.Context) {

	var cName = c.Request.FormValue("Computer")
	var cUser = c.Request.FormValue("User")
	var cFile = c.Request.FormValue("Screen")

	o, err := os.Create(paths.GenerateMonitoringFile(cName, cUser))

	if err != nil {
		fmt.Println("Error1 = " + err.Error())
		return
	}

	defer o.Close()

	_, err2 := o.WriteString(cFile)

	if err2 != nil {
		fmt.Println("Error1 = " + err2.Error())
		return
	}

	return

}
