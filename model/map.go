package model

import (
	"gorm.io/gorm"
)

type Map struct {
	gorm.Model
	Image    string
	Sections []Section `gorm:"foreignKey:MapRefer"`
}
