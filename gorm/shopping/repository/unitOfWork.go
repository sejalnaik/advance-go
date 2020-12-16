package repository

import "github.com/jinzhu/gorm"

type UnitOfWork struct {
	DB        *gorm.DB
	Committed bool
	ReadOnly  bool
}
