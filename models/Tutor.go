package models

import "gorm.io/gorm"

type Tutor struct {
	gorm.Model
	UserID         uint     `gorm:"uniqueIndex"`
	User           User     `gorm:"constraint:OnDelete:CASCADE;"`
	Qualifications string   `gorm:"not null"`
	Expertise      string   `gorm:"not null"`
	Bio            string   `gorm:"type:text"`
	Experience     int      `gorm:"not null"`
	Availability   string   `gorm:"not null"`
	Rating         float32  `gorm:"default:0"`
	Courses        []Course `gorm:"foreignKey:TutorID"`
}
