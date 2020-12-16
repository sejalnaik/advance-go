package crud

import (
	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/many-to-many/model"
)

func LoadPeople(db *gorm.DB) {
	AddPerson(db, model.Person{Name: "sejal"})
	AddPerson(db, model.Person{Name: "rachel"})
	AddPerson(db, model.Person{Name: "ross"})
}

func LoadLanguages(db *gorm.DB) {
	AddLanguage(db, model.Language{Name: "Hindi"})
	AddLanguage(db, model.Language{Name: "English"})
	AddLanguage(db, model.Language{Name: "Marathi"})
}

func AddPerson(db *gorm.DB, person model.Person) {
	db.Debug().Create(&person)
}

func AddLanguage(db *gorm.DB, language model.Language) {
	db.Debug().Create(&language)
}

func GetPerson(db *gorm.DB, id int) model.Person {
	tempPerson := model.Person{}
	db.Debug().First(&tempPerson, id)
	return tempPerson
}

func GetLanguage(db *gorm.DB, id int) model.Language {
	tempLanguage := model.Language{}
	db.Debug().First(&tempLanguage, id)
	return tempLanguage
}

func UpdatePerson(db *gorm.DB, person model.Person) {
	db.Debug().Model(&person).Update(&person)
}

func UpdateLanguage(db *gorm.DB, language model.Language) {
	db.Debug().Model(&language).Update(&language)
}

func GetLanguagesFromPeople(db *gorm.DB, id int) []model.Language {
	tempPerson := model.Person{}
	db.Debug().Preload("Languages").First(&tempPerson, id)
	return tempPerson.Languages
}

func GetPeopleFromLanguages(db *gorm.DB, id int) []model.Person {
	templanguage := model.Language{}
	db.Debug().Preload("Person").First(&templanguage, id)
	return templanguage.Persons
}
