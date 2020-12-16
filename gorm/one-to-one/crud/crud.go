package crud

import (
	"fmt"

	"github.com/sejalnaik/advance-go/gorm/one-to-one/user"

	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/one-to-one/profile"
)

func AddUsers(db *gorm.DB) {
	AddUser(db, user.User{Name: "sejal"})
	AddUser(db, user.User{Name: "ross"})
	AddUser(db, user.User{Name: "rachel"})
	AddUser(db, user.User{Name: "phoebe"})
}

func AddUser(db *gorm.DB, user user.User) {
	db.Create(&user)
	isAdded := db.NewRecord(user)
	fmt.Println(isAdded)
	if isAdded == false {
		fmt.Print("Added!!")
	} else {
		fmt.Println("Not Added!!")
	}
}

func AddProfiles(db *gorm.DB) {
	AddProfile(db, profile.Profile{UserID: 1, ProfileName: "sejalnaik"})
	AddProfile(db, profile.Profile{UserID: 2, ProfileName: "rossgeller"})
	AddProfile(db, profile.Profile{UserID: 3, ProfileName: "rachelgreen"})
	AddProfile(db, profile.Profile{UserID: 4, ProfileName: "phoebebuffay"})
}

func AddProfile(db *gorm.DB, profile profile.Profile) {
	db.Create(&profile)
	isAdded := db.NewRecord(profile)
	fmt.Println(isAdded)
	if isAdded == false {
		fmt.Print("Added!!")
	} else {
		fmt.Println("Not Added!!")
	}
}

func GetAllprofilesWithUserID(db *gorm.DB, id int) profile.Profile {
	tempProfile := profile.Profile{}
	db.First(&user.User{}, id).Related(&tempProfile)
	return tempProfile
}
