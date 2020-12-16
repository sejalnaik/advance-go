package rachel

import "github.com/jinzhu/gorm"

type Rachel struct {
	gorm.Model
	Name          string
	AccountNumber int
	Balance       float64
}
