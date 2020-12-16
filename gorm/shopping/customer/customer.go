package customer

import (
	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/shopping/order"
)

type Customer struct {
	gorm.Model
	Name   string
	Orders []order.Order
}
