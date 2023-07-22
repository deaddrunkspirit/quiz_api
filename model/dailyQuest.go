package model

import (
	"gorm.io/gorm"
)

type DailyQuest struct {
	gorm.Model
	Description          string
	TargetScore          int
	Reward               int
	RewardType           string
	DailyQuestProgresses []DailyQuestProgress `gorm:"foreignKey:DailyQuestRefer"`
}
