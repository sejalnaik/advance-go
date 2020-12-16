package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/shopping/crud"
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
	/*order1 := order.Order{Qunatity: 2, Cost: 100.00, CustomerID: 1}
	order2 := order.Order{Qunatity: 3, Cost: 150.00, CustomerID: 1}
	orders1 := []order.Order{order1, order2}
	customerSejal := customer.Customer{Name: "sejal", Orders: orders1}
	crud.AddCustomerAndOrders(db, customerSejal)
	order3 := order.Order{Qunatity: 4, Cost: 800.00, CustomerID: 2}
	order4 := order.Order{Qunatity: 5, Cost: 950.00, CustomerID: 2}
	orders2 := []order.Order{order3, order4}
	customerRachel := customer.Customer{Name: "rachel", Orders: orders2}
	crud.AddCustomerAndOrders(db, customerRachel)*/

	//Get Cutomer by ID
	/*customerSejal := crud.GetCustomer(db, 1)
	fmt.Println(customerSejal.Name)
	fmt.Println(len(customerSejal.Orders))
	fmt.Println(customerSejal.Orders[0])
	fmt.Println(customerSejal.Orders[1])*/

	//Get orders by cutomer ID
	//fmt.Println(crud.GetOrdersWithCustomerID(db, 1))

	//Update customer info
	/*customerSejal := customer.Customer{Name: "sejal naik", Model: gorm.Model{ID: 1}}
	crud.UpdateCustomerInfo(db, customerSejal)*/

	//Update order info through customer ID
	/*customerSejal := crud.GetCustomer(db, 6)
	customerSejal.Orders[0].Qunatity = 100
	customerSejal.Orders[0].Cost = 10000
	crud.UpdateOrderInfoThroughCustomerID(db, customerSejal)*/

	/*customerSejal := crud.GetCustomer(db, 6)
	orderNew := order.Order{Qunatity: 50, Cost: 20000.00}
	customerSejal.Orders = append(customerSejal.Orders, orderNew)
	crud.UpdateOrderInfoThroughCustomerID(db, customerSejal)*/

	/*customerSejal := crud.GetCustomer(db, 6)
	customerSejal.Orders = customerSejal.Orders[0:2]
	fmt.Println(len(customerSejal.Orders))
	crud.UpdateOrderInfoThroughCustomerID(db, customerSejal)*/

	//delete customer(soft)
	/*customerRachel := crud.GetCustomer(db, 11)
	crud.DeleteCustomer(db, customerRachel)*/

	//count rows scoped
	/*fmt.Println("Scoped count of customer:", crud.ScopedCountRows(db, customer.Customer{}))
	fmt.Println("Scoped count of order:", crud.ScopedCountRows(db, order.Order{}))*/

	//count rows unscoped
	/*fmt.Println("UnScoped count of customer:", crud.UnScopedCountRows(db, customer.Customer{}))
	fmt.Println("UnScoped count of order:", crud.UnScopedCountRows(db, order.Order{}))*/

	//delete customer(delete)
	customerRachel := crud.GetCustomer(db, 13)
	crud.HardDeleteCustomer(db, customerRachel)
}
