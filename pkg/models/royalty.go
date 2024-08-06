package models

import "gorm.io/gorm"

type Royalty struct {
	gorm.Model
	TokenID          uint `gorm:"index:idx_royalty_token_id"`
	UserID           uint `gorm:"index:idx_royalty_user_id"`
	Amount           float64
	DistributionDate string
}
