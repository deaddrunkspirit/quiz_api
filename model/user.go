package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password              string                `json:"-"`
	Username              string                `gorm:"size:255;not null;unique" json:"username"`
	Email                 string                `gorm:"size:255;not null;" json:"email"`
	Avatar                string                `json:"avatar"`
	LastVisit             time.Time             `json:"last_visit"`
	DayStreak             int                   `json:"day_streak"`
	AchievementProgresses []AchievementProgress `gorm:"foreignKey:UserRefer" json:"achievement_progresses"`
	DailyQuestProgresses  []DailyQuestProgress  `gorm:"foreignKey:UserRefer" json:"daily_quest_progresses"`
	Wallet                []Wallet              `gorm:"foreignKey:UserRefer" json:"wallet"`
	UserAnswers           []UserAnswer          `gorm:"foreignKey:UserRefer" json:"user_answers"`
	LevelScores           []LevelScore          `gorm:"foreignKey:UserRefer" json:"level_scores"`
}
