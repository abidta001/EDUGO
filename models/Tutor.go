package models

import "gorm.io/gorm"

type Tutor struct {
	gorm.Model
	UserID         uint     `gorm:"uniqueIndex"` // Foreign key to User
	User           User     `gorm:"constraint:OnDelete:CASCADE;"`
	Qualifications string   `gorm:"not null"`           // Degrees, certifications, etc.
	Expertise      string   `gorm:"not null"`           // Subjects or skills taught
	Bio            string   `gorm:"type:text"`          // Short introduction
	Experience     int      `gorm:"not null"`           // Years of teaching experience
	Availability   string   `gorm:"not null"`           // Available days/hours
	Rating         float32  `gorm:"default:0"`          // Average rating from students
	Courses        []Course `gorm:"foreignKey:TutorID"` // Related courses
}
