package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// declaring a model

type Employee struct {
	gorm.Model
	code       string
	name       string
	department string
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrating the schema
	db.AutoMigrate(&Employee{})

	// create
	db.Create(&Employee{code: "1234", name: "Mahesh Wosti", department: "Infrastructure"})

	// read
	var emp Employee
	db.First(&emp, 1)                // find integer with primary key
	db.First(&emp, "code=?", "1234") // find product with code

	// Update- update the emp department
	db.Model(&emp).Update("department", "Admin")

	db.Model(&emp).Updates(Employee{name: "hari", code: "4567"}) // updates- update multiple field

}
