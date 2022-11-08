package models

import (
	"gorm.io/gorm"
)

type UUDBridge struct {
	gorm.Model
	UID  uint
	UDID uint
}
