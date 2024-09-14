package model

import (
	"gorm.io/gorm"
)

type Program struct {
	gorm.Model
	DisplayName string       `json:"displayName"`
	ProgramDay  []ProgramDay `gorm:"foreignKey:ProgramID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"programDay"`
}
