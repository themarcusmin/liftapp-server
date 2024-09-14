package model

type ExerciseWorkoutJunction struct {
	ID            uint    `gorm:"primaryKey"`
	SetNumber     *uint8  `json:"prescribedSet,omitempty"` // ordering
	PrescribedRep *uint8  `json:"prescribedRep,omitempty"`
	RestTime      *uint16 `json:"prescribedRestTime,omitempty"`
	ExerciseID    uint
	WorkoutID     *uint
	LogEntries    []LogEntry `gorm:"foreignKey:JunctionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"logEntries"`
}
