package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	MusicID      uint   `gorm:"index:idx_token_music_id"`
	TokenAddress string `gorm:"unique"`
}
