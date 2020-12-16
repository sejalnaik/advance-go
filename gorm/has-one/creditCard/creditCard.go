package creditCard

import "github.com/jinzhu/gorm"

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
