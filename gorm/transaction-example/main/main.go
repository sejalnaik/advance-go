package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/transaction-example/crud"
	"github.com/sejalnaik/advance-go/gorm/transaction-example/rachel"
	"github.com/sejalnaik/advance-go/gorm/transaction-example/sejal"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&sejal.Sejal{}, &rachel.Rachel{})

	//Load data
	//crud.LoadRachel(db)
	//crud.LoadSejal(db)

	//Transaction
	tempSejal := crud.GetSejal(db, 1)
	tempRachel := crud.GetRachel(db, 1)
	tempSejal.Balance -= 100.00
	tempRachel.Balance += 100.00
	if err := crud.TransferTransaction(db, tempSejal, tempRachel); err != nil {
		fmt.Println("Transaction succesful")
	} else {
		fmt.Println("Transaction failed")
	}
}
