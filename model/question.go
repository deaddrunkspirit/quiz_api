package model

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Text            int
	Options         []int
	CorrectAnswerId int
	Image           string
	QuestionType    string
	Duration        int
	Difficulty      int
	UserAnswers     []UserAnswer `gorm:"foreignKey:QuestionRefer"`
	QuestionPack    QuestionPack `gorm:"foreignKey:QuestionPackRefer"`
}
