package state

import (
	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/associations/city"
)

type State struct {
	gorm.Model
	Name   string
	Cities []city.City
}
