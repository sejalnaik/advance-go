package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/custom-model/country"
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
	/*tempCountryID := "0ff09404-8042-41da-ac12-6fc853c9e818"
	tempCountry := crud.GetCountry(db, tempCountryID)
	fmt.Println(tempCountry)*/

	//update country
	/*tempCountryID := "0ff09404-8042-41da-ac12-6fc853c9e818"
	tempCountry := crud.GetCountry(db, tempCountryID)
	tempCountry.Name = "United states"
	crud.UpdateCountry(db, tempCountry)
	fmt.Println(tempCountry)*/

	/*tempCountryID := "0ff09404-8042-41da-ac12-6fc853c9e818"
	tempCountry := crud.GetCountry(db, tempCountryID)
	tempCountry.States = append(tempCountry.States, state.State{Name: "Kansas"})
	crud.UpdateCountry(db, tempCountry)
	fmt.Println(tempCountry)*/

	//delete country(soft)
	/*tempCountryID := "0ff09404-8042-41da-ac12-6fc853c9e818"
	tempCountry := crud.GetCountry(db, tempCountryID)
	crud.DeleteCountry(db, tempCountry)*/

	//delete country (hard)
	/*tempCountryID := "0ff09404-8042-41da-ac12-6fc853c9e818"
	tempCountry := crud.GetCountry(db, tempCountryID)
	crud.HardDeleteCountry(db, tempCountry)*/

}
