package crud

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/shopping/customer"
	"github.com/sejalnaik/advance-go/gorm/shopping/order"
)

func AddCustomerAndOrders(db *gorm.DB, customer customer.Customer) {
	db.Debug().Create(&customer)
}

func GetCustomer(db *gorm.DB, id int) customer.Customer {
	tempCustomer := customer.Customer{}
	db.Debug().Preload("Orders").First(&tempCustomer, id)
	return tempCustomer
}

func GetOrdersWithCustomerID(db *gorm.DB, id int) []order.Order {
	tempOrders := []order.Order{}
	db.Debug().First(&customer.Customer{}, id).Related(&tempOrders)
	return tempOrders
}

func UpdateCustomerInfo(db *gorm.DB, customerNew customer.Customer) {
	db.Debug().Model(&customerNew).Update(map[string]interface{}{"name": customerNew.Name})
}

func UpdateOrderInfoThroughCustomerID(db *gorm.DB, customer customer.Customer) {
	db.Debug().Model(&customer).Update(&customer)
}

func ScopedCountRows(db *gorm.DB, table interface{}) int {
	var count int
	fmt.Println(&table)
	fmt.Println(*(&table))
	db.Debug().Model(table).Count(&count)
	return count
}

func UnScopedCountRows(db *gorm.DB, table interface{}) int {
	var count int
	fmt.Println(&table)
	fmt.Println(*(&table))
	db.Debug().Model(table).Unscoped().Count(&count)
	return count
}

func DeleteCustomer(db *gorm.DB, customer customer.Customer) {
	if len(customer.Orders) != 0 {
		fmt.Println("length is not 0!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		for _, order := range customer.Orders {
			db.Debug().Model(&order).Delete(&order)
		}
	}
	db.Debug().Delete(&customer)
}

func HardDeleteCustomer(db *gorm.DB, customer customer.Customer) {
	db.Debug().Unscoped().Delete(&customer)
}
