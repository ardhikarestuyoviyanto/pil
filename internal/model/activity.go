package model

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	ID        int    `gorm:"primaryKey, AUTO_INCREMENT"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	LinkYt    string `json:"link_yt"`
	LinkDrive string `json:"link_drive"`
}

type APIResponseActivity struct {
	ID        int    `json:"id"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	LinkYt    string `json:"link_yt"`
	LinkDrive string `json:"link_drive"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
