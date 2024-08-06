package models

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	UserId     uint   `gorm:"index:idx_artists_user_id"`
	ArtistName string `gorm:"unique"`
	Bio        string
}
