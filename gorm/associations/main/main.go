package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/associations/city"
	"github.com/sejalnaik/advance-go/gorm/associations/crud"
	"github.com/sejalnaik/advance-go/gorm/associations/state"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&state.State{}, &city.City{})
	db.Model(&city.City{}).AddForeignKey("state_id", "states(id)", "CASCADE", "CASCADE")

	//load states
	//crud.LoadStates(db)

	//get cities by associations
	tempCities := crud.StateAssociations(db, 1)
	fmt.Println("length of slice :", len(tempCities))
	for i, city := range tempCities {
		fmt.Println("City:", i, ":", city)
	}

	//append cities to state
	/*addcities := []city.City{{Name: "satara"}, {Name: "Navi mumbai"}}
	crud.StateAddAssociations(db, 1, addcities)*/

	//replace associations
	/*replaceCities := []city.City{{Name: "empty1"}, {Name: "empty2"}}
	crud.StateReplaceAssociations(db, 2, replaceCities)*/

	//count associations
	//fmt.Println("Count:", crud.StateCountAssociations(db, 1))

	//delete associations
	/*tempCities := crud.StateAssociations(db, 2)
	crud.StatDeleteAssociations(db, 2, tempCities)*/
}
