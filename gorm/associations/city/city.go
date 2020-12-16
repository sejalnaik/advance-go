package city

import "github.com/jinzhu/gorm"

type City struct {
	gorm.Model
	Name    string
	StateID uint
}
