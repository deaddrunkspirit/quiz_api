package model

import (
	"gorm.io/gorm"
)

type AchievementProgress struct {
	gorm.Model
	Score       int
	IsCompleted bool
	User        User        `gorm:"foregnKey:UserRefer"`
	Achievement Achievement `gorm:"foreignKey:AchievementRefer"`
}
