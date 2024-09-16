package handler

import (
	"net/http"

	logrus "github.com/sirupsen/logrus"

	gdatabase "liftapp/database"
	gmodel "liftapp/database/model"

	"liftapp/app/database/model"
)

// CreateLog handles jobs for controller.CreateLog
func CreateLog(userIDAuth uint64, log model.Log) (httpResponse gmodel.HTTPResponse, httpStatusCode int) {
	db := gdatabase.GetDB()
	user := model.User{}
	logFinal := model.Log{}

	// check if the user have an existing profile
	if err := db.Where("id_auth = ?", userIDAuth).First(&user).Error; err != nil {
		httpResponse.Message = "no user profile found"
		httpStatusCode = http.StatusForbidden
		return
	}

	logFinal.UserID = log.UserID
	logFinal.EventAt = log.EventAt
	logFinal.LogExercise = log.LogExercise
	logFinal.ProgramDayID = log.ProgramDayID

	tx := db.Begin()
	if err := tx.Create(&logFinal).Error; err != nil {
		tx.Rollback()
		logrus.WithError(err).Error("error code: 1211")
		httpResponse.Message = "internal server error"
		httpStatusCode = http.StatusInternalServerError
		return
	}
	tx.Commit()

	httpResponse.Message = logFinal
	httpStatusCode = http.StatusCreated
	return
}
