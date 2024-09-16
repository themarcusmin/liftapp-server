package model

type Format struct {
	ID          uint   `gorm:"primaryKey"`
	DisplayName string `gorm:"not null;unique" json:"displayName"`
}
