package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// update single column
func singleColumn(conn *gorm.DB) {
	result := conn.Model(&User{}).Where("Name = ?", "brandon").Update("Email", "brandonichami630@gmail.com")
	if result.Error != nil {
		log.Fatal(result.Error)
	} else {
		fmt.Println("single column update successful..")
	}
}

// multiple column update
func multipleColumn(conn *gorm.DB) {
	result := conn.Model(&User{}).Where("Name = ?", "john").Updates(User{Name: "jonah", Email: "jonah@gmail.com"})
	if result.Error != nil {
		log.Fatal(result.Error)
	} else {
		fmt.Println("multiple column update successful..")
	}
}

// delete record
func delete(conn *gorm.DB) {
	result := conn.Delete(&User{}, 1)
	if result.Error != nil {
		log.Fatal(result.Error)
	} else {
		fmt.Println("deleted successfully..")
	}
}
