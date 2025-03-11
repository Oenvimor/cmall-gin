package model

import (
	"github.com/jinzhu/gorm"
)

// Product 商品模型
type Product struct {
	gorm.Model
	Name          string
	CategoryID    int
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
}
