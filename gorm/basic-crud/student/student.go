package student

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name string  `gorm:"type:varchar(100)"`
	Age  int     `gorm:"type:int"`
	Cgpa float32 `gorm:"type:decimal(3,1)"`
}
