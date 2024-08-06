package models

import "gorm.io/gorm"

type Music struct {
	gorm.Model
	ArtistID          uint `gorm:"index:idx_music_artist_id"`
	Title             string
	Description       string
	RoyaltyPercentage float64
}
