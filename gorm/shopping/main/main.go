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

	//Add customer with orders
	/*uow := repository.UnitOfWork{DB: db, Committed: false, ReadOnly: false}
	order1 := order.Order{Qunatity: 2, Cost: 100.00, CustomerID: 1}
	order2 := order.Order{Qunatity: 3, Cost: 150.00, CustomerID: 1}
	orders1 := []order.Order{order1, order2}
	customerSejal := customer.Customer{Name: "sejal", Orders: orders1}
	repository.AddCustomerAndOrders(uow, customerSejal)
	order3 := order.Order{Qunatity: 4, Cost: 800.00, CustomerID: 2}
	order4 := order.Order{Qunatity: 5, Cost: 950.00, CustomerID: 2}
	orders2 := []order.Order{order3, order4}
	customerRachel := customer.Customer{Name: "rachel", Orders: orders2}
	repository.AddCustomerAndOrders(uow, customerRachel)*/

	//Get Cutomer by ID
	/*uow := repository.UnitOfWork{DB: db, Committed: false, ReadOnly: true}
	customerSejal := repository.GetCustomer(uow, 14)
	fmt.Println(customerSejal.Name)
	fmt.Println(len(customerSejal.Orders))
	fmt.Println(customerSejal.Orders[0])
	fmt.Println(customerSejal.Orders[1])*/

	//Get orders by cutomer ID
	/*uow := repository.UnitOfWork{DB: db, Committed: false, ReadOnly: true}
	fmt.Println(repository.GetOrdersWithCustomerID(uow, 14))*/

	//Update customer info
	/*uow := repository.UnitOfWork{DB: db, Committed: false, ReadOnly: false}
	customerSejal := customer.Customer{Name: "sejal naik", Model: gorm.Model{ID: 14}}
	repository.UpdateCustomerInfo(uow, customerSejal)*/

	//Update order info through customer ID
	/*uow := repository.UnitOfWork{DB: db, Committed: false, ReadOnly: true}
	customerSejal := repository.GetCustomer(uow, 14)
	uow = repository.UnitOfWork{DB: db, Committed: false, ReadOnly: false}
	customerSejal.Orders[0].Qunatity = 100
	customerSejal.Orders[0].Cost = 10000
	repository.UpdateOrderInfoThroughCustomerID(uow, customerSejal)*/

	/*uow := repository.UnitOfWork{DB: db, Committed: false, ReadOnly: false}
	customerSejal := repository.GetCustomer(uow, 14)
	orderNew := order.Order{Qunatity: 50, Cost: 20000.00}
	customerSejal.Orders = append(customerSejal.Orders, orderNew)
	repository.UpdateOrderInfoThroughCustomerID(uow, customerSejal)*/

	//delete customer(soft)
	/*uow := repository.UnitOfWork{DB: db, Committed: false, ReadOnly: true}
	customerRachel := repository.GetCustomer(uow, 14)
	uow = repository.UnitOfWork{DB: db, Committed: false, ReadOnly: false}
	repository.DeleteCustomer(uow, customerRachel)*/

	//count rows scoped
	/*uow := repository.UnitOfWork{DB: db, Committed: false, ReadOnly: true}
	fmt.Println("Scoped count of customer:", repository.ScopedCountRows(uow, customer.Customer{}))
	fmt.Println("Scoped count of order:", repository.ScopedCountRows(uow, order.Order{}))*/

	//count rows unscoped
	/*uow := repository.UnitOfWork{DB: db, Committed: false, ReadOnly: true}
	fmt.Println("UnScoped count of customer:", repository.UnScopedCountRows(uow, customer.Customer{}))
	fmt.Println("UnScoped count of order:", repository.UnScopedCountRows(uow, order.Order{}))*/

	//delete customer(hard)
	/*uow := repository.UnitOfWork{DB: db, Committed: false, ReadOnly: true}
	customerRachel := repository.GetCustomer(uow, 16)
	uow = repository.UnitOfWork{DB: db, Committed: false, ReadOnly: false}
	repository.HardDeleteCustomer(uow, customerRachel)*/
}
