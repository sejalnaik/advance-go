package sejal

import "github.com/jinzhu/gorm"

type Sejal struct {
	gorm.Model
	Name          string
	AccountNumber int
	Balance       float64
}
