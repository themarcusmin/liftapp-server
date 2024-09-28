package handler

import (
	"net/http"

	gdatabase "liftapp/database"
	gmodel "liftapp/database/model"

	log "github.com/sirupsen/logrus"
)

type GetProgramDaysResponse struct {
	ID          uint   `gorm:"column:id" json:"ID"`
	DisplayName string `gorm:"column:display_name" json:"displayName"`
	ProgramID   uint   `gorm:"column:program_id" json:"programID"`
}

// GetProgramDays handles jobs for controller.GetProgramDays
func GetProgramDays(userIDAuth uint64) (httpResponse gmodel.HTTPResponse, httpStatusCode int) {
	db := gdatabase.GetDB()

	programDaysResponse := []GetProgramDaysResponse{}

	if err := db.
		Table("program_days AS pD").
		Select("pD.id, pD.display_name, pD.program_id").
		Joins("JOIN programs AS p ON pD.program_id = p.id").
		Where("p.user_id = ?", userIDAuth).
		Find(&programDaysResponse).
		Error; err != nil {
		log.WithError(err)
		httpResponse.Message = "internal server error"
		httpStatusCode = http.StatusInternalServerError
		return
	}

	if len(programDaysResponse) == 0 {
		httpResponse.Message = "no program found"
		httpStatusCode = http.StatusNotFound
		return
	}

	httpResponse.Message = programDaysResponse
	httpStatusCode = http.StatusOK
	return
}
