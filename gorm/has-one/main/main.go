package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/has-one/creditCard"
	"github.com/sejalnaik/advance-go/gorm/has-one/crud"
	"github.com/sejalnaik/advance-go/gorm/has-one/user"
)

func main() {
	// connect to database
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	// create table
	creditCardEmpty := creditCard.CreditCard{}
	userEmpty := user.User{}
	db.AutoMigrate(&creditCardEmpty, &userEmpty)

	//load credit cards
	//crud.AddCreditCards(db)

	//add users
	/*tempCreditCard := creditCard.CreditCard{}
	db.First(&tempCreditCard)
	crud.AddUser(db, tempCreditCard)*/

	//get credit card info from user
	fmt.Println(crud.GetCreditCardWithUserID(db, 1))
}
