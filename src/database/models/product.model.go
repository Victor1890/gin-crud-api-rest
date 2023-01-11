package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"name" gorm:"type:text"`
	Brand string  `json:"brand" gorm:"type:text"`
	Size  uint64  `json:"size" gorm:"type:real"`
	Price float64 `json:"price" gorm:"not null;type:float"`
	Files []File  `json:"files"`
}
