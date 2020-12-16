package profile

import (
	"github.com/sejalnaik/advance-go/gorm/one-to-one/user"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	UserID      int
	User        user.User `gorm:"foreignkey:UserID"`
	ProfileName string
}
