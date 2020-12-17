package repository

import (
	"github.com/sejalnaik/advance-go/gorm/shopping/customer"
	"github.com/sejalnaik/advance-go/gorm/shopping/order"
)

type gormRepository struct {
}

func NewRepository() *gormRepository {
	return &gormRepository{}
}

func (*gormRepository) AddCustomerAndOrders(uow *UnitOfWork, customer customer.Customer) error {
	db := uow.DB

	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		return err
	}

	if err := db.Debug().Create(&customer).Error; err != nil {
		return err
	}

	return nil
}

func (*gormRepository) GetCustomer(uow *UnitOfWork, id int) (customer.Customer, error) {
	db := uow.DB
	tempCustomer := customer.Customer{}

	if err := db.Debug().Preload("Orders").First(&tempCustomer, id).Error; err != nil {
		return tempCustomer, err
	}

	return tempCustomer, nil
}

func (*gormRepository) GetOrdersWithCustomerID(uow *UnitOfWork, id int) ([]order.Order, error) {
	db := uow.DB
	tempOrders := []order.Order{}

	if err := db.Debug().First(&customer.Customer{}, id).Related(&tempOrders).Error; err != nil {
		return tempOrders, err
	}

	return tempOrders, nil
}

func (*gormRepository) UpdateCustomerInfo(uow *UnitOfWork, customerNew customer.Customer) error {

	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		return err
	}

	if err := db.Debug().Model(&customerNew).Update(map[string]interface{}{"name": customerNew.Name}).Error; err != nil {
		return err
	}

	return nil
}

func (*gormRepository) UpdateOrderInfoThroughCustomerID(uow *UnitOfWork, customer customer.Customer) error {

	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		return err
	}

	if err := db.Debug().Model(&customer).Update(&customer).Error; err != nil {
		return err
	}

	return nil
}

func (*gormRepository) ScopedCountRows(uow *UnitOfWork, table interface{}) (int, error) {

	db := uow.DB
	var count int

	if err := db.Debug().Model(table).Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

func (*gormRepository) UnScopedCountRows(uow *UnitOfWork, table interface{}) (int, error) {

	db := uow.DB
	var count int

	if err := db.Debug().Model(table).Unscoped().Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

func (*gormRepository) DeleteCustomer(uow *UnitOfWork, customer customer.Customer) error {

	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		return err
	}

	if len(customer.Orders) != 0 {
		for _, order := range customer.Orders {
			if err := db.Debug().Model(&order).Delete(&order).Error; err != nil {
				return err
			}
		}
	}

	if err := db.Debug().Delete(&customer).Error; err != nil {
		return err
	}

	return nil
}

func (*gormRepository) HardDeleteCustomer(uow *UnitOfWork, customer customer.Customer) error {

	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		return err
	}
	if err := db.Debug().Unscoped().Delete(&customer).Error; err != nil {
		return err
	}

	return nil
}
