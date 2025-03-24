package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string   `gorm:"not null" json:"title"`
	Description string   `gorm:"type:text" json:"description"`
	CategoryID  uint     `gorm:"not null" json:"category_id"`
	Category    Category `gorm:"constraint:OnDelete:CASCADE;" json:"category"`
	Price       float64  `json:"price"`
	TutorID     uint     `gorm:"not null" json:"tutor_id"`
	Tutor       Tutor    `gorm:"constraint:OnDelete:CASCADE;" json:"tutor"`
}
