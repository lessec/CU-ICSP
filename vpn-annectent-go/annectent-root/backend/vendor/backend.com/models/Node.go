package models

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	Name     string `json:"name"`
	IP       string `json:"ip"`
	Domain   string `json:"domain"`
	Describe string `json:"describe"`
}
