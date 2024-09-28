package controller

import (
	"reflect"

	"github.com/gin-gonic/gin"

	grenderer "liftapp/lib/renderer"

	"liftapp/app/handler"
)

/*
GetProgramDays - GET /programDays
*/
func GetProgramDays(c *gin.Context) {
	userIDAuth := c.GetUint64("authID")

	resp, statusCode := handler.GetProgramDays(userIDAuth)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)
}
