package crud

import (
	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/transaction-example/rachel"
	"github.com/sejalnaik/advance-go/gorm/transaction-example/sejal"
)

func LoadSejal(db *gorm.DB) {
	db.Debug().Create(&sejal.Sejal{Name: "sejal", AccountNumber: 1234, Balance: 200.00})
}

func LoadRachel(db *gorm.DB) {
	db.Debug().Create(&rachel.Rachel{Name: "rachel", AccountNumber: 5678, Balance: 0.00})
}

func GetSejal(db *gorm.DB, id int) sejal.Sejal {
	tempSejal := sejal.Sejal{}
	db.Debug().First(&tempSejal, id)
	return tempSejal
}

func GetRachel(db *gorm.DB, id int) rachel.Rachel {
	tempRachel := rachel.Rachel{}
	db.Debug().First(&tempRachel, id)
	return tempRachel
}

func TransferTransaction(db *gorm.DB, sejal sejal.Sejal, rachel rachel.Rachel) error {
	transaction := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	if err := transaction.Error; err != nil {
		return err
	}

	if err := transaction.Debug().Model(&sejal).Update(&sejal).Error; err != nil {
		transaction.Rollback()
		return err
	}
	//panic("panic!!!!!!!!!!!!!!")

	if err := transaction.Debug().Model(&rachel).Update(&rachel).Error; err != nil {
		transaction.Rollback()
		return err
	}

	return transaction.Commit().Error
}
