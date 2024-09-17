package controller

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"

	grenderer "liftapp/lib/renderer"

	"liftapp/app/handler"
)

/*
GetRecent1RM - GET /logEntry
*/
func GetRecent1RM(c *gin.Context) {
	userIDAuth := c.GetUint64("authID")
	exerciseID := strings.TrimSpace(c.Params.ByName("exercise_id"))

	resp, statusCode := handler.GetRecent1RM(userIDAuth, exerciseID)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)
}
