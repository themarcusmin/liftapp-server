package handler

import (
	"math"
	"net/http"
	"time"

	logrus "github.com/sirupsen/logrus"

	gdatabase "liftapp/database"
	gmodel "liftapp/database/model"

	"liftapp/app/database/model"
)

// CreateLog handles jobs for controller.CreateLog
func CreateLog(userIDAuth uint64, log model.Log) (httpResponse gmodel.HTTPResponse, httpStatusCode int) {
	db := gdatabase.GetDB()
	logFinal := model.Log{}

	logFinal.UserID = userIDAuth
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

// Response type for GetLogs
type LogResponse struct {
	ID            uint            `json:"ID"`
	EventAt       time.Time       `json:"event_at"`
	Sets          uint            `json:"sets"`
	Volume        uint            `json:"volume"`
	MusclesWorked map[string]uint `json:"musclesWorked"`
}

/*
GetLogs handles jobs for controller.GetLogs returning a list of entries
Each entry includes:
- ID: a unique identifier for the log
- EventAt: the timestamp of the event in ISO 8601 format.
- Sets: the total number of sets recorded in each log
- Volume: the total volume worked in each log
- MusclesWorked: a map where each key is a muscle name and the value is the number of sets targeting that muscle. Sets of non-primary muscles worked are halved
*/
func GetLogs(userIDAuth uint64, startTime string, endTime string) (httpResponse gmodel.HTTPResponse, httpStatusCode int) {
	db := gdatabase.GetDB()
	logsResponse := []LogResponse{}

	if err := db.Table("logs").
		Select(`
			logs.id,
			logs.event_at,
			COUNT(DISTINCT log_entries.id) as sets,
			SUM(COALESCE(log_entries.reps, 0) * COALESCE(log_entries.weight, 0)) as volume
		`).
		Joins("JOIN log_exercises ON log_exercises.log_id = logs.id").
		Joins("JOIN log_entries ON log_entries.log_exercise_id = log_exercises.id").
		Where("logs.user_id = ? AND logs.event_at BETWEEN ? AND ?", userIDAuth, startTime, endTime).
		Group("logs.id").
		Scan(&logsResponse).
		Error; err != nil {
		logrus.WithError(err)
		httpResponse.Message = "internal server error"
		httpStatusCode = http.StatusInternalServerError
		return
	}

	// Fetch muscles worked for each logResponse
	for i, logResponse := range logsResponse {
		var muscles []struct {
			DisplayName string
			IsPrimary   bool
			SetCount    uint
		}

		if err := db.Table("muscles").
			Select("muscles.display_name, exercise_muscles.is_primary, COUNT(DISTINCT log_entries.id) as set_count").
			Joins("JOIN exercise_muscles ON exercise_muscles.muscle_id = muscles.id").
			Joins("JOIN log_exercises ON log_exercises.exercise_id = exercise_muscles.exercise_id").
			Joins("JOIN log_entries ON log_entries.log_exercise_id = log_exercises.id").
			Joins("JOIN logs ON logs.id = log_exercises.log_id").
			Where("logs.id = ?", logResponse.ID).
			Group("muscles.id, exercise_muscles.is_primary").
			Scan(&muscles).Error; err != nil {
			logrus.WithError(err)
			httpResponse.Message = "internal server error"
			httpStatusCode = http.StatusInternalServerError
			return
		}

		logsResponse[i].MusclesWorked = make(map[string]uint)
		for _, muscle := range muscles {
			count := muscle.SetCount
			if !muscle.IsPrimary {
				count = uint(math.Ceil(float64(count) / 2))

			}
			logsResponse[i].MusclesWorked[muscle.DisplayName] = count
		}
	}

	httpResponse.Message = logsResponse
	httpStatusCode = http.StatusOK
	return
}
