package country

import (
	"github.com/sejalnaik/advance-go/gorm/custom-model/model"
	"github.com/sejalnaik/advance-go/gorm/custom-model/state"
)

type Country struct {
	model.Model
	Name   string
	States []state.State
}
