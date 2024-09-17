package handler

import (
	"net/http"

	logrus "github.com/sirupsen/logrus"

	gdatabase "liftapp/database"
	gmodel "liftapp/database/model"
)

// Response type for GetRecent1RM
type Recent1RMResponse struct {
	ID      uint    `json:"id"`
	Reps    uint8   `json:"reps"`
	Weight  float64 `json:"weight"`
	OneRm   float64 `json:"oneRm"`
	EventAt string  `json:"eventAt"`
}

/*
GetRecent1RM handles jobs for controller.GetRecent1RM returning the log entry having 1RM from the past 5 logs
1RM is calculated using epley formula: (weight * (1 + reps / 30.0))
*/
func GetRecent1RM(userIDAuth uint64, exerciseID string) (httpResponse gmodel.HTTPResponse, httpStatusCode int) {
	db := gdatabase.GetDB()
	var recent1RMResponse Recent1RMResponse

	if err := db.Table("log_entries").
		Select("log_entries.id, log_entries.reps, log_entries.weight, ROUND(log_entries.weight * (1 + log_entries.reps / 30.0), 2) AS one_rm, log_entries.event_at").
		Joins("JOIN log_exercises ON log_exercises.id = log_entries.log_exercise_id").
		Joins("JOIN logs ON logs.id = log_exercises.log_id").
		Where("log_exercises.exercise_id = ?", exerciseID).
		Where("logs.user_id = ?", userIDAuth).
		Order("logs.event_at DESC").
		Limit(5).
		Order("one_rm DESC").
		First(&recent1RMResponse).
		Error; err != nil {
		logrus.WithError(err)
		httpResponse.Message = "internal server error"
		httpStatusCode = http.StatusInternalServerError
		return
	}

	httpResponse.Message = recent1RMResponse
	httpStatusCode = http.StatusOK
	return

}
