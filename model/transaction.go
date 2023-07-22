package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Direction   string
	Description string
	Amount      int
	Wallet      Wallet `gorm:"foreignKey:WalletRefer"`
}
