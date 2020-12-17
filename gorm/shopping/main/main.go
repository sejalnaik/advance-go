package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/shopping/customer"
	"github.com/sejalnaik/advance-go/gorm/shopping/order"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&customer.Customer{}, &order.Order{})
	db.Model(&order.Order{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "RESTRICT")

	//Add customer with orders***************************************************************************
	/*uow := repository.NewUnitOfWork(db, false)
	repo := repository.NewRepository()
	order1 := order.Order{Qunatity: 2, Cost: 100.00, CustomerID: 1}
	order2 := order.Order{Qunatity: 3, Cost: 150.00, CustomerID: 1}
	orders1 := []order.Order{order1, order2}
	customerSejal := customer.Customer{Name: "sejal", Orders: orders1}
	if err := repo.AddCustomerAndOrders(uow, customerSejal); err != nil {
		uow.Complete()
	} else {
		uow.Commit()
	}*/

	/*uow := repository.NewUnitOfWork(db, false)
	repo := repository.NewRepository()
	order3 := order.Order{Qunatity: 4, Cost: 800.00, CustomerID: 2}
	order4 := order.Order{Qunatity: 5, Cost: 950.00, CustomerID: 2}
	orders2 := []order.Order{order3, order4}
	customerRachel := customer.Customer{Name: "rachel", Orders: orders2}
	if err := repo.AddCustomerAndOrders(uow, customerRachel); err != nil {
		uow.Complete()
	} else {
		uow.Commit()
	}*/

	//Get Cutomer by ID****************************************************************************
	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	if customerSejal, err := repo.GetCustomer(uow, 18); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(customerSejal.Name)
		fmt.Println(len(customerSejal.Orders))
		fmt.Println(customerSejal.Orders[0])
		fmt.Println(customerSejal.Orders[1])
	}*/

	//Get orders by cutomer ID**************************************************************************
	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	if tempOrders, err := repo.GetOrdersWithCustomerID(uow, 18); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tempOrders)
	}*/

	//Update customer info******************************************************************************
	/*uow := repository.NewUnitOfWork(db, false)
	repo := repository.NewRepository()
	customerSejal := customer.Customer{Name: "sejal naik", Model: gorm.Model{ID: 18}}
	if err := repo.UpdateCustomerInfo(uow, customerSejal); err != nil {
		uow.Complete()
	} else {
		uow.Commit()
	}*/

	//Update order info through customer ID***************************************************************
	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	customerSejal, err := repo.GetCustomer(uow, 18)
	if err != nil {
		fmt.Println(err)
	} else {
		uow = repository.NewUnitOfWork(db, false)
		customerSejal.Orders[0].Qunatity = 100
		customerSejal.Orders[0].Cost = 10000
		if err := repo.UpdateOrderInfoThroughCustomerID(uow, customerSejal); err != nil {
			uow.Complete()
		} else {
			uow.Commit()
		}
	}*/

	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	customerSejal, err := repo.GetCustomer(uow, 18)
	if err != nil {
		fmt.Println(err)
	} else {
		uow = repository.NewUnitOfWork(db, false)
		orderNew := order.Order{Qunatity: 50, Cost: 20000.00}
		customerSejal.Orders = append(customerSejal.Orders, orderNew)
		if err := repo.UpdateOrderInfoThroughCustomerID(uow, customerSejal); err != nil {
			uow.Complete()
		} else {
			uow.Commit()
		}
	}*/

	//delete customer(soft)*****************************************************************************
	/*uow := repository.NewUnitOfWork(db, false)
	repo := repository.NewRepository()
	customerRachel, err := repo.GetCustomer(uow, 19)
	if err != nil {
		fmt.Println(err)
	} else {
		uow = repository.NewUnitOfWork(db, false)
		if err := repo.DeleteCustomer(uow, customerRachel); err != nil {
			uow.Complete()
		} else {
			uow.Commit()
		}
	}*/

	//count rows scoped************************************************************************************
	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	if count, err := repo.ScopedCountRows(uow, customer.Customer{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Scoped count of customer:", count)
	}*/

	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	if count, err := repo.ScopedCountRows(uow, order.Order{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Scoped count of order:", count)
	}*/

	//count rows unscoped*********************************************************************************
	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	if count, err := repo.UnScopedCountRows(uow, customer.Customer{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("UnScoped count of customer:", count)
	}*/

	/*uow := repository.NewUnitOfWork(db, true)
	repo := repository.NewRepository()
	if count, err := repo.UnScopedCountRows(uow, order.Order{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("UnScoped count of order:", count)
	}*/

	//delete customer(hard)******************************************************************************
	/*uow := repository.NewUnitOfWork(db, false)
	repo := repository.NewRepository()
	customerSejal, err := repo.GetCustomer(uow, 18)
	if err != nil {
		fmt.Println(err)
	} else {
		uow = repository.NewUnitOfWork(db, false)
		if err := repo.HardDeleteCustomer(uow, customerSejal); err != nil {
			uow.Complete()
		} else {
			uow.Commit()
		}
	}*/
}
