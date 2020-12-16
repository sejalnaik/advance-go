package crud

import (
	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/custom-model/country"
	"github.com/sejalnaik/advance-go/gorm/custom-model/state"
)

func LoadCountries(db *gorm.DB) {
	AddCountry(db, country.Country{Name: "India",
		States: []state.State{
			{Name: "Maharashtra"},
			{Name: "Gujurat"},
		},
	})
	AddCountry(db, country.Country{Name: "USA",
		States: []state.State{
			{Name: "New York"},
			{Name: "Ney Jersey"},
		},
	})
}

func AddCountry(db *gorm.DB, country country.Country) {
	db.Debug().Create(&country)
}

func GetCountry(db *gorm.DB, id string) country.Country {
	tempCountry := country.Country{}
	db.Debug().Preload("States").Where("id = ?", id).First(&tempCountry)
	return tempCountry
}

func UpdateCountry(db *gorm.DB, country country.Country) {
	db.Debug().Model(&country).Update(&country)
}

func DeleteCountry(db *gorm.DB, country country.Country) {

	if len(country.States) != 0 {
		for _, state := range country.States {
			db.Debug().Model(&state).Delete(&state)
		}
	}
	db.Debug().Delete(&country)
}
func HardDeleteCountry(db *gorm.DB, country country.Country) {
	db.Debug().Unscoped().Delete(&country)
}
