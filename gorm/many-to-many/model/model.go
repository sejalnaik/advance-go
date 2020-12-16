package model

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:person_languages;"`
}

type Language struct {
	gorm.Model
	Name    string
	Persons []Person `gorm:"many2many:person_languages;"`
}
