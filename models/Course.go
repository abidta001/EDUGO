package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
	Price       float64
	TutorID     uint  `gorm:"not null"` 
	Tutor       Tutor `gorm:"constraint:OnDelete:CASCADE;"`
}
