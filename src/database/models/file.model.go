package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name      string `gorm:"type:text"`
	MimeType  string `gorm:"type:text"`
	Extension string `gorm:"type:text"`
	Url       string `gorm:"type:text"`
	Size      int64  `gorm:"type:int4"`
	ProductId uint   `json:"product_id"`
}
