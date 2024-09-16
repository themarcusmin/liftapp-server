package controller

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	grenderer "liftapp/lib/renderer"

	"liftapp/app/database/model"
	"liftapp/app/handler"
)

/*
CreateLog - POST /logs
*/
func CreateLog(c *gin.Context) {
	userIDAuth := c.GetUint64("authID")
	log := model.Log{}

	// bind JSON
	if err := c.ShouldBindJSON(&log); err != nil {
		grenderer.Render(c, gin.H{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	resp, statusCode := handler.CreateLog(userIDAuth, log)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)

}

/*
GetLog - GET /logs
*/
func GetLogs(c *gin.Context) {
	userIDAuth := c.GetUint64("authID")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")

	resp, statusCode := handler.GetLogs(userIDAuth, startTime, endTime)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)
}
