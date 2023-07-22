package model

import (
	"gorm.io/gorm"
)

type LevelScore struct {
	gorm.Model
	IsCup bool
	Score int
	Level Level `gorm:"foreignKey:LevelRefer"`
	User  User  `gorm:"foreignKey:UserRefer"`
}
