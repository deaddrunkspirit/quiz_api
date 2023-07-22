package model

import (
	"gorm.io/gorm"
)

type UserAnswer struct {
	gorm.Model
	Answer   string
	User     User     `gorm:"foreignKey:UserRefer"`
	Question Question `gorm:"foreignKey:QuestionRefer"`
}
