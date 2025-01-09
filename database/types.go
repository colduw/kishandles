package database

import (
	"gorm.io/gorm"
)

type (
	BskyHandle struct {
		gorm.Model
		Handle string `gorm:"column:handle;unique"`
		DID    string `gorm:"column:did;unique"`
	}

	DiscordHandle struct {
		gorm.Model
		UserName string `gorm:"column:user_name;unique"`
		DHCode   string `gorm:"column:dh_code;unique"`
	}
)
