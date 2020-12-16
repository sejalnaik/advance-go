package crud

import (
	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/associations/city"
	"github.com/sejalnaik/advance-go/gorm/associations/state"
)

func LoadStates(db *gorm.DB) {
	AddState(db, state.State{Name: "maharashtra",
		Cities: []city.City{
			{Name: "mumbai"},
			{Name: "pune"},
		},
	})
	AddState(db, state.State{Name: "gujurat",
		Cities: []city.City{
			{Name: "surat"},
			{Name: "ahmedabad"},
		},
	})
}

func AddState(db *gorm.DB, state state.State) {
	db.Debug().Create(&state)
}

func StateAssociations(db *gorm.DB, id int) []city.City {
	tempCities := []city.City{}
	db.Debug().First(&state.State{}, id).Association("Cities").Find(&tempCities)
	return tempCities
}

func StateAddAssociations(db *gorm.DB, id int, cities []city.City) {
	db.Debug().First(&state.State{}, id).Association("Cities").Append(cities)
}

func StateReplaceAssociations(db *gorm.DB, id int, cities []city.City) {
	db.Debug().First(&state.State{}, id).Association("Cities").Replace(cities)
}

func StateCountAssociations(db *gorm.DB, id int) int {
	return db.Debug().First(&state.State{}, id).Association("Cities").Count()
}

func StatDeleteAssociations(db *gorm.DB, id int, cities []city.City) {
	db.Debug().First(&state.State{}, id).Association("Cities").Delete(cities)
}
