package state

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sejalnaik/advance-go/gorm/custom-model/model"
)

type State struct {
	model.Model
	Name      string
	CountryID uuid.UUID `gorm:"type:varchar(36)"`
}
