package order

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Qunatity   int
	Cost       float64
	CustomerID uint
}
