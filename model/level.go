package model

import (
	"gorm.io/gorm"
)

type Level struct {
	gorm.Model
	TargetScore int
	LevelScores []LevelScore `gorm:"foreignKey:LevelRefer"`
	Section     Section      `gorm:"foreignKey:SectionRefer"`
}
