package controller

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	grenderer "liftapp/lib/renderer"

	model "liftapp/app/database/model"
	"liftapp/app/handler"
)

/*
CreateProgram - POST /programs
*/
func CreateProgram(c *gin.Context) {
	userIDAuth := c.GetUint64("authID")
	req := model.Program{}

	// bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		grenderer.Render(c, gin.H{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	resp, statusCode := handler.CreateProgram(userIDAuth, req)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)
}
