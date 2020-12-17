package model

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Name   string
	Orders []Order
}
