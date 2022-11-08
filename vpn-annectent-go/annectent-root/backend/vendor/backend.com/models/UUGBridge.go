package models

import (
	"gorm.io/gorm"
)

type UUGBridge struct {
	gorm.Model
	UID  uint
	UGID uint
}
