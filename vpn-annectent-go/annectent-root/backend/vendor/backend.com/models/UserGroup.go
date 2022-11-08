package models

import (
	"gorm.io/gorm"
)

type UserGroup struct {
	gorm.Model
	Name          string
	Describe      string
	GrantedServer string
}

type SerializedUserGroup struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Describe      string `json:"describe"`
	GrantedServer string `json:"grantedServer"`
}
