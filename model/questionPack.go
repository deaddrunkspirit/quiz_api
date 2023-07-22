package model

import (
	"gorm.io/gorm"
)

type QuestionPack struct {
	gorm.Model
	Questions []Question `gorm:"foreignKey:QuestionPackRefer"`
}
