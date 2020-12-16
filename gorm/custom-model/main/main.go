package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/custom-model/country"
	"github.com/sejalnaik/advance-go/gorm/custom-model/crud"
	"github.com/sejalnaik/advance-go/gorm/custom-model/state"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/gorm_custom_model?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&country.Country{}, &state.State{})
	db.Model(&state.State{}).AddForeignKey("country_id", "countries(id)", "CASCADE", "CASCADE")

	//load countries
	//crud.LoadCountries(db)

	//get country by id
	/*uuidAsUUID, _ := uuid.FromString("0ff09404-8042-41da-ac12-6fc853c9e818")
	tempCountry := crud.GetCountry(db, uuidAsUUID)
	fmt.Println(tempCountry)*/

	tempCountry := crud.GetFirstCountry(db)
	tempCountryID := tempCountry.ID
	gotCountry := crud.GetCountry(db, tempCountryID)
	fmt.Println(gotCountry)
}
