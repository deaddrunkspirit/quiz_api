package model

import (
	"gorm.io/gorm"
)

type DailyQuestProgress struct {
	gorm.Model
	Score       int
	IsCompleted bool
	DailyQuest  DailyQuest `gorm:"foreignKey:DailyQuestRefer"`
	User        User       `gorm:"foreignKey:UserRefer"`
}
