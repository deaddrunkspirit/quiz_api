package model

import (
	"gorm.io/gorm"
)

type Achievement struct {
	gorm.Model
	Title                 string
	Levels                []int
	Image                 string
	Description           string
	AchievementProgresses []Achievement `gorm:"foreignKey:AchievementRefer"`
}
