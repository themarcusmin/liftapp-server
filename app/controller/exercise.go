package controller

import (
	"reflect"

	"github.com/gin-gonic/gin"

	grenderer "liftapp/lib/renderer"

	"liftapp/app/handler"
)

/*
GetExercisesByMuscleGroup - GET /exercises
*/
func GetExercisesByMuscleGroup(c *gin.Context) {
	resp, statusCode := handler.GetExercisesByMuscleGroup()

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)
}
