package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/one-to-one/crud"
	"github.com/sejalnaik/advance-go/gorm/one-to-one/profile"
	"github.com/sejalnaik/advance-go/gorm/one-to-one/user"
)

func main() {
	// connect to database
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	// create table
	userEmpty := user.User{}
	profileEmpty := profile.Profile{}
	db.AutoMigrate(&userEmpty, &profileEmpty)

	// load users
	//crud.AddUsers(db)

	//load profiles
	//crud.AddProfiles(db)

	//get profile of given user id*****************not getting user************************
	fmt.Println(crud.GetAllprofilesWithUserID(db, 3))
}
