package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/many-to-many/crud"
	"github.com/sejalnaik/advance-go/gorm/many-to-many/model"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&model.Person{}, &model.Language{})

	//load people and languages
	/*crud.LoadPeople(db)
	crud.LoadLanguages(db)*/

	//give languages to sejal
	/*sejalPerson := crud.GetPerson(db, 1)
	language1 := crud.GetLanguage(db, 1)
	language2 := crud.GetLanguage(db, 2)

	sejalPerson.Languages = append(sejalPerson.Languages, language1)
	sejalPerson.Languages = append(sejalPerson.Languages, language2)

	crud.UpdatePerson(db, sejalPerson)*/

	//give languages to rachel
	/*sejalRachel := crud.GetPerson(db, 2)
	language1 := crud.GetLanguage(db, 2)
	language2 := crud.GetLanguage(db, 3)

	sejalRachel.Languages = append(sejalRachel.Languages, language1)
	sejalRachel.Languages = append(sejalRachel.Languages, language2)

	crud.UpdatePerson(db, sejalRachel)*/

	//get langauges of person
	/*tempLanguages := crud.GetLanguagesFromPeople(db, 1)
	fmt.Println("length of slice :", len(tempLanguages))
	for i, language := range tempLanguages {
		fmt.Println("Language:", i, ":", language)
	}*/

	//get people of language
	tempPeople := crud.GetPeopleFromLanguages(db, 1)
	fmt.Println("length of slice :", len(tempPeople))
	for i, person := range tempPeople {
		fmt.Println("Language:", i, ":", person)
	}
}
