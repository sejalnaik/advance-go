package crud

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/has-one/creditCard"
	"github.com/sejalnaik/advance-go/gorm/has-one/user"
)

func AddCreditCards(db *gorm.DB) {
	AddCreditCard(db, creditCard.CreditCard{Number: "12345678"})
	AddCreditCard(db, creditCard.CreditCard{Number: "38353855"})
	AddCreditCard(db, creditCard.CreditCard{Number: "48548674"})
	AddCreditCard(db, creditCard.CreditCard{Number: "28478272"})
}

func AddCreditCard(db *gorm.DB, creditCard creditCard.CreditCard) {
	db.Create(&creditCard)
	isAdded := db.NewRecord(creditCard)
	fmt.Println(isAdded)
	if isAdded == false {
		fmt.Print("Added!!")
	} else {
		fmt.Println("Not Added!!")
	}
}

func AddUser(db *gorm.DB, creditCard creditCard.CreditCard) {
	tempUser := user.User{CreditCard: creditCard}
	db.Create(&tempUser)
	isAdded := db.NewRecord(tempUser)
	fmt.Println(isAdded)
	if isAdded == false {
		fmt.Print("Added!!")
	} else {
		fmt.Println("Not Added!!")
	}
}

func GetCreditCardWithUserID(db *gorm.DB, id int) creditCard.CreditCard {
	tempCreditCard := creditCard.CreditCard{}
	db.First(&user.User{}).Related(&tempCreditCard)
	return tempCreditCard
}
