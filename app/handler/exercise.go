package handler

import (
	"net/http"

	gdatabase "liftapp/database"
	gmodel "liftapp/database/model"

	"liftapp/app/database/model"
)

type GetExercisesResponse struct {
	MuscleGroups []MuscleGroupResponse `json:"muscleGroups"`
}

type MuscleGroupResponse struct {
	Muscle    string             `json:"muscle"`
	Exercises []ExerciseResponse `json:"exercises"`
}

type ExerciseResponse struct {
	ExerciseID   uint                  `json:"exerciseId"`
	ExerciseName string                `json:"exerciseName"`
	SubExercises []SubExerciseResponse `json:"subExercises,omitempty"`
}

type SubExerciseResponse struct {
	ExerciseID   uint   `json:"exerciseId"`
	ExerciseName string `json:"exerciseName"`
}

/*
GetExercisesByMuscleGroup handles jobs for controller.GetExercisesByMuscleGroup returning a list of exercises grouped by muscles
*/
func GetExercisesByMuscleGroup() (httpResponse gmodel.HTTPResponse, httpStatusCode int) {
	db := gdatabase.GetDB()
	getExercisesResponse := GetExercisesResponse{}

	muscles := []model.Muscle{}

	if err := db.Preload("Exercise").Find(&muscles).Error; err != nil {
		httpResponse.Message = "internal server error"
		httpStatusCode = http.StatusInternalServerError
		return
	}

	for _, muscle := range muscles {
		muscleGroupResponse := MuscleGroupResponse{
			Muscle: muscle.DisplayName,
		}

		// Track parent exercises
		exerciseMap := make(map[uint]*ExerciseResponse)

		// Populate exercises for the muscle
		for _, exercise := range muscle.Exercise {
			// Only consider parent exercises
			if exercise.ParentID == nil {
				exerciseResponse := &ExerciseResponse{
					ExerciseID:   exercise.ID,
					ExerciseName: exercise.DisplayName,
					// SubExercises: []SubExerciseResponse{}, // Initialize with an empty list
				}

				// Add sub-exercises
				for _, subExercise := range muscle.Exercise {
					// Check if the exercise is a child of the current parent
					if subExercise.ParentID != nil && *subExercise.ParentID == exercise.ID {
						subExerciseResponse := SubExerciseResponse{
							ExerciseID:   subExercise.ID,
							ExerciseName: subExercise.DisplayName,
						}
						exerciseResponse.SubExercises = append(exerciseResponse.SubExercises, subExerciseResponse)
					}
				}

				// Store the parent exercise in the map
				exerciseMap[exercise.ID] = exerciseResponse
			}
		}

		// Convert the map to a slice for the final response
		for _, exerciseResponse := range exerciseMap {
			muscleGroupResponse.Exercises = append(muscleGroupResponse.Exercises, *exerciseResponse)
		}

		getExercisesResponse.MuscleGroups = append(getExercisesResponse.MuscleGroups, muscleGroupResponse)
	}

	httpResponse.Message = getExercisesResponse
	httpStatusCode = http.StatusOK
	return
}
