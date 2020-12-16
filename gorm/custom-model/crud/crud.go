package crud

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
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

func GetCountry(db *gorm.DB, id uuid.UUID) country.Country {
	tempCountry := country.Country{}
	db.Debug().Preload("States").Where("id = ?", id.String()).First(&tempCountry)
	return tempCountry
}

func GetFirstCountry(db *gorm.DB) country.Country {
	tempCountry := country.Country{}
	db.Debug().Preload("States").First(&tempCountry)
	return tempCountry
}
