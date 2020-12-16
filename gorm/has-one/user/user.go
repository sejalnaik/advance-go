package user

import (
	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/has-one/creditCard"
)

type User struct {
	gorm.Model
	CreditCard creditCard.CreditCard
}
