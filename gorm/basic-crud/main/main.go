package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sejalnaik/advance-go/gorm/basic-crud/student"
)

func main() {
	// connect to database
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	// create table
	studentEmpty := student.Student{}
	db.AutoMigrate(&studentEmpty)

	// add new student
	/*newStudent := student.Student{Name: "rachel", Age: 36, Cgpa: 9.8}
	crud.AddStudent(db, newStudent)
	defer db.Close()*/

	//query first
	/*student := crud.QueryFirst(db)
	fmt.Println(student)*/

	//query find
	/*students := crud.QueryFind(db)
	fmt.Println(students)*/

	// query first with id
	/*student := crud.QueryFirstWithID(db, 4)
	fmt.Println(student)*/

	//update student
	//crud.UpdateStudent(db, 1, "sejal naik", 30, 9.0)

	//delete student
	//crud.DeleteStudent(db, 4)
}
