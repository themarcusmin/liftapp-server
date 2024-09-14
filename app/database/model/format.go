package model

type Format struct {
	ID          uint   `gorm:"primaryKey"`
	DisplayName string `json:"displayName"`
}
