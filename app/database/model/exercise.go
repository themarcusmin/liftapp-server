package model

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	DisplayName string `json:"displayName,omitempty"`
}
