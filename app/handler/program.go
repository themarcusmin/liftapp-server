package handler

import (
	"net/http"

	gdatabase "liftapp/database"
	gmodel "liftapp/database/model"

	"liftapp/app/database/model"

	log "github.com/sirupsen/logrus"
)

// CreateProgram handles jobs for controller.CreateProgram
func CreateProgram(userIDAuth uint64, req model.Program) (httpResponse gmodel.HTTPResponse, httpStatusCode int) {
	db := gdatabase.GetDB()
	user := model.User{}

	// does the user have an existing profile
	if err := db.Where("id_auth = ?", userIDAuth).First(&user).Error; err != nil {
		httpResponse.Message = "no user profile found"
		httpStatusCode = http.StatusForbidden
		return
	}

	tx := db.Begin()

	// create program
	program := model.Program{
		DisplayName: req.DisplayName,
		UserID:      userIDAuth,
	}

	if err := tx.Create(&program).Error; err != nil {
		tx.Rollback()
		log.WithError(err)
		httpResponse.Message = "Failed to create program"
		httpStatusCode = http.StatusInternalServerError
		return
	}

	// create programDay, programExercises, and programEntries
	for _, dayReq := range req.ProgramDay {
		programDay := model.ProgramDay{
			ProgramID:   program.ID,
			DisplayName: dayReq.DisplayName,
		}

		if err := tx.Create(&programDay).Error; err != nil {
			tx.Rollback()
			log.WithError(err)
			httpResponse.Message = "Failed to create program day"
			httpStatusCode = http.StatusInternalServerError
			return
		}

		for _, exerciseReq := range dayReq.ProgramExercise {
			programExercise := model.ProgramExercise{
				ProgramDayID: programDay.ID,
				ExerciseID:   exerciseReq.ExerciseID,
				RestTime:     exerciseReq.RestTime,
			}

			if err := tx.Create(&programExercise).Error; err != nil {
				tx.Rollback()
				log.WithError(err)
				httpResponse.Message = "Failed to create program exercise"
				httpStatusCode = http.StatusInternalServerError
				return
			}

			for _, entryReq := range exerciseReq.ProgramEntry {
				programEntry := model.ProgramEntry{
					ProgramExerciseID:         programExercise.ID,
					SetNumber:                 entryReq.SetNumber,
					PrescribedReps:            entryReq.PrescribedReps,
					PrescribedOneRmPrecentage: entryReq.PrescribedOneRmPrecentage,
					PrescribedDuration:        entryReq.PrescribedDuration,
				}

				if err := tx.Create(&programEntry).Error; err != nil {
					tx.Rollback()
					log.WithError(err)
					httpResponse.Message = "Failed to create program entry"
					httpStatusCode = http.StatusInternalServerError
					return
				}
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.WithError(err)
		httpResponse.Message = "Failed to commit transaction"
		httpStatusCode = http.StatusInternalServerError
		return
	}

	httpResponse.Message = program
	httpStatusCode = http.StatusCreated
	return
}
