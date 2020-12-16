package repository

import (
	"github.com/sejalnaik/advance-go/gorm/shopping/customer"
	"github.com/sejalnaik/advance-go/gorm/shopping/order"
)

func AddCustomerAndOrders(uow UnitOfWork, customer customer.Customer) error {

	transaction := uow.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	if err := transaction.Error; err != nil {
		return err
	}

	if err := transaction.Debug().Create(&customer).Error; err != nil {
		transaction.Rollback()
		return err
	}
	uow.Committed = true

	return transaction.Commit().Error
}

func GetCustomer(uow UnitOfWork, id int) customer.Customer {
	tempCustomer := customer.Customer{}
	uow.DB.Debug().Preload("Orders").First(&tempCustomer, id)
	uow.Committed = true
	return tempCustomer
}

func GetOrdersWithCustomerID(uow UnitOfWork, id int) []order.Order {
	tempOrders := []order.Order{}
	uow.DB.Debug().First(&customer.Customer{}, id).Related(&tempOrders)
	return tempOrders
}

func UpdateCustomerInfo(uow UnitOfWork, customerNew customer.Customer) error {

	transaction := uow.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	if err := transaction.Error; err != nil {
		return err
	}

	if err := transaction.Debug().Model(&customerNew).Update(map[string]interface{}{"name": customerNew.Name}).Error; err != nil {
		transaction.Rollback()
		return err
	}
	uow.Committed = true

	return transaction.Commit().Error
}

func UpdateOrderInfoThroughCustomerID(uow UnitOfWork, customer customer.Customer) error {

	transaction := uow.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	if err := transaction.Error; err != nil {
		return err
	}

	if err := transaction.Debug().Model(&customer).Update(&customer).Error; err != nil {
		transaction.Rollback()
		return err
	}
	uow.Committed = true

	return transaction.Commit().Error
}

func ScopedCountRows(uow UnitOfWork, table interface{}) int {
	var count int
	uow.DB.Debug().Model(table).Count(&count)
	return count
}

func UnScopedCountRows(uow UnitOfWork, table interface{}) int {
	var count int
	uow.DB.Debug().Model(table).Unscoped().Count(&count)
	return count
}

func DeleteCustomer(uow UnitOfWork, customer customer.Customer) error {

	transaction := uow.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	if err := transaction.Error; err != nil {
		return err
	}

	if len(customer.Orders) != 0 {
		for _, order := range customer.Orders {
			if err := transaction.Debug().Model(&order).Delete(&order).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	if err := transaction.Debug().Delete(&customer).Error; err != nil {
		transaction.Rollback()
		return err
	}

	uow.Committed = true

	return transaction.Commit().Error
}

func HardDeleteCustomer(uow UnitOfWork, customer customer.Customer) error {

	transaction := uow.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	if err := transaction.Error; err != nil {
		return err
	}

	if err := transaction.Debug().Unscoped().Delete(&customer).Error; err != nil {
		transaction.Rollback()
		return err
	}
	uow.Committed = true

	return transaction.Commit().Error
}
