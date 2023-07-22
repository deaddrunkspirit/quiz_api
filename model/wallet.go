package model

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Amount       int
	WalletType   string
	User         User          `gorm:"foreignKey:UserRefer"`
	Transactions []Transaction `gorm:"foreignKey:WalletRefer"`
}
