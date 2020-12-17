package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/shopping/model"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&model.Customer{}, &model.Order{})
	db.Model(&model.Order{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "RESTRICT")

	//Add customer with orders***************************************************************************
	/*uow := repository.NewUnitOfWork(db, false)
	repo := repository.NewRepository()
	order1 := model.Order{Qunatity: 2, Cost: 100.00, CustomerID: 1}
	order2 := model.Order{Qunatity: 3, Cost: 150.00, CustomerID: 1}
	orders1 := []model.Order{order1, order2}
	customerSejal := model.Customer{Name: "sejal", Orders: orders1}
	if err := repo.Add(uow, &customerSejal); err != nil {
		uow.Complete()
	} else {
		uow.Commit()
	}*/

	/*uow := repository.NewUnitOfWork(db, false)
	repo := repository.NewRepository()
	order1 := model.Order{Qunatity: 4, Cost: 500.00, CustomerID: 2}
	order2 := model.Order{Qunatity: 5, Cost: 750.00, CustomerID: 2}
	orders1 := []model.Order{order1, order2}
	customerRachel := model.Customer{Name: "rachel", Orders: orders1}
	if err := repo.Add(uow, &customerRachel); err != nil {
		uow.Complete()
	} else {
		uow.Commit()
	}*/

	//Get Cutomer by ID****************************************************************************
	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	tempCustomer := model.Customer{Name: "sejal naik", Model: gorm.Model{ID: 21}}
	if err := repo.Get(uow, &tempCustomer); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tempCustomer.Name)
		fmt.Println(len(tempCustomer.Orders))
		fmt.Println(tempCustomer.Orders[0])
		fmt.Println(tempCustomer.Orders[1])
	}*/

	//Get orders by cutomer ID**************************************************************************
	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	tempCustomer := model.Customer{Model: gorm.Model{ID: 21}}
	if tempOrders, err := repo.GetOrdersWithCustomerID(uow, &tempCustomer); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tempOrders)
	}*/

	//Update customer info******************************************************************************
	/*uow := repository.NewUnitOfWork(db, false)
	repo := repository.NewRepository()
	customerSejal := model.Customer{Name: "sejal naik", Model: gorm.Model{ID: 21}}
	if err := repo.Update(uow, &customerSejal); err != nil {
		uow.Complete()
	} else {
		uow.Commit()
	}*/

	//Update order info through customer ID***************************************************************
	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	tempCustomer := model.Customer{Model: gorm.Model{ID: 21}}
	if err := repo.Get(uow, &tempCustomer); err != nil {
		fmt.Println(err)
	} else {
		uow = repository.NewUnitOfWork(db, false)
		tempCustomer.Orders[0].Qunatity = 100
		tempCustomer.Orders[0].Cost = 10000
		if err := repo.Update(uow, &tempCustomer); err != nil {
			uow.Complete()
		} else {
			uow.Commit()
		}
	}*/

	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	tempCustomer := model.Customer{Model: gorm.Model{ID: 21}}
	if err := repo.Get(uow, &tempCustomer); err != nil {
		fmt.Println(err)
	} else {
		uow = repository.NewUnitOfWork(db, false)
		orderNew := model.Order{Qunatity: 50, Cost: 20000.00}
		tempCustomer.Orders = append(tempCustomer.Orders, orderNew)
		if err := repo.Update(uow, &tempCustomer); err != nil {
			uow.Complete()
		} else {
			uow.Commit()
		}
	}*/

	//delete customer(soft)*****************************************************************************
	/*uow := repository.NewUnitOfWork(db, false)
	repo := repository.NewRepository()
	tempCustomer := model.Customer{Model: gorm.Model{ID: 22}}
	if err := repo.Get(uow, &tempCustomer); err != nil {
		fmt.Println(err)
	} else {
		uow = repository.NewUnitOfWork(db, false)
		if err := repo.DeleteCustomer(uow, tempCustomer); err != nil {
			uow.Complete()
		} else {
			uow.Commit()
		}
	}*/

	//count rows scoped************************************************************************************
	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	if count, err := repo.ScopedCountRows(uow, &model.Customer{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Scoped count of customer:", count)
	}*/

	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	if count, err := repo.ScopedCountRows(uow, &model.Order{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Scoped count of order:", count)
	}*/

	//count rows unscoped*********************************************************************************
	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	if count, err := repo.UnScopedCountRows(uow, &model.Customer{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("UnScoped count of customer:", count)
	}*/

	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	if count, err := repo.UnScopedCountRows(uow, &model.Order{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("UnScoped count of order:", count)
	}*/

	//delete customer(hard)******************************************************************************
	/*uow := repository.NewUnitOfWork(db, false)
	repo := repository.NewRepository()
	tempCustomer := model.Customer{Model: gorm.Model{ID: 22}}
	if err != nil {
		fmt.Println(err)
	} else {
		uow = repository.NewUnitOfWork(db, false)
		if err := repo.HardDelete(uow, &tempCustomer); err != nil {
			uow.Complete()
		} else {
			uow.Commit()
		}
	}*/
}
