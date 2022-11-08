package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
	Type  string
}

type UserType string

const (
	Admin   string = "Admin"
	Manager string = "Manager"
	Staff   string = "Staff"
)

type SerializedUser struct {
	ID      uint         `json:"id"`
	Name    string       `json:"name"`
	Email   string       `json:"email"`
	Type    string       `json:"type"`
	Devices []UserDevice `json:"devices"`
	Groups  []UserGroup  `json:"groups"`
}
