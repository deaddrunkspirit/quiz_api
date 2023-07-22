package model

import (
	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	Image  string
	Map    Map     `gorm:"foreignKey:MapRefer"`
	Levels []Level `gorm:"foreignKey:SectionRefer"`
}
