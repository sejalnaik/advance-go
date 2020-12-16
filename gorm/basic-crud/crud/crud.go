package crud

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/advance-go/gorm/basic-crud/student"
)

func AddStudent(db *gorm.DB, student student.Student) {
	db.Create(&student)
	isAdded := db.NewRecord(student)
	fmt.Println(isAdded)
	if isAdded == false {
		fmt.Print("Added!!")
	} else {
		fmt.Println("Not Added!!")
	}
}

func QueryFirst(db *gorm.DB) student.Student {
	student := student.Student{}
	db.First(&student)
	return student
}

func QueryFind(db *gorm.DB) []student.Student {
	students := []student.Student{}
	db.Find(&students)
	return students
}

func QueryFirstWithID(db *gorm.DB, id int) student.Student {
	student := student.Student{}
	db.First(&student, id)
	return student
}

func UpdateStudent(db *gorm.DB, id int, name string, age int, cgpa float32) {
	/*tempStudent := student.Student{}
	db.First(&tempStudent, id)
	fmt.Print("Before update:", tempStudent)
	tempStudent.Name = name
	tempStudent.Age = age
	tempStudent.Cgpa = cgpa
	db.Save(&tempStudent)
	fmt.Println("After Update:", tempStudent)*/
	db.First(&student.Student{}, id).Updates(student.Student{Name: "sejal", Age: 26, Cgpa: 7.8})
}

func DeleteStudent(db *gorm.DB, id int) {
	tempStudent := student.Student{}
	db.First(&tempStudent, id).Delete(&tempStudent)
}
