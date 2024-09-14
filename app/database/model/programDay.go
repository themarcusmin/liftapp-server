package model

import (
	"gorm.io/gorm"
)

type ProgramDay struct {
	gorm.Model
	DisplayName string `json:"displayName"`
	ProgramID   uint
}
