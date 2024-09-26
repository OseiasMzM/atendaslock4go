package handler

import (
	"atendaslock4go/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetText(c *gin.Context) {
	logs.AddText("aew")
	c.String(http.StatusOK, "")
}

func GetText(c *gin.Context) {
	c.String(http.StatusOK, logs.ToJson())
}

func PostTest(c *gin.Context) {
	println(c.Request.PostFormValue("a"))
	c.String(http.StatusOK, "Successor")
}
