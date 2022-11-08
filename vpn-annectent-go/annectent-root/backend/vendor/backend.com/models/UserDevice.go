package models

import "gorm.io/gorm"

type UserDevice struct {
	gorm.Model
	Name       string
	PrivateKey string
	PublicKey  string
	QRCode     string
	ConfigFile string
	IP         string
}

type SerializedUserDevice struct {
	ID         uint   `json:"id"`
	IP         string `json:"ip"`
	Name       string `json:"name"`
	PublicKey  string `json:"publicKey"`
	QRCode     string `json:"qr_code"`
	ConfigFile string `json:"configFile"`
}
