package models

import "gorm.io/gorm"

type Loketdua struct {
	gorm.Model
	Seq int `json:"ant"`
	Stat string `json:"status"`
}