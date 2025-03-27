package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// selecting by id
func selectById(conn *gorm.DB) {
	var user User
	result := conn.First(&user, 2)
	if result.Error != nil {
		log.Fatal(result.Error)
	} else {
		fmt.Println("select by id", user)
	}
}

// select by condition
func getByEmail(conn *gorm.DB) {
	var user User
	conn.First(&user, "Email = ?", "brandonichami@gmail.com")
	fmt.Println(user)
}

// select all users
func getAll(conn *gorm.DB) {
	var users []User
	conn.Find(&users)
	fmt.Println(users)
}

// selecting specific columns
func getColumn(conn *gorm.DB) {
	var users []User
	conn.Select("Name").Find(&users)
	fmt.Println(users)
}

// order clause
func getOrderBy(conn *gorm.DB) {
	var users []User
	conn.Order("Name desc").Find(&users)
	fmt.Println(users)
}

// limit and offset
func pagination(conn *gorm.DB) {
	var users []User
	conn.Limit(2).Offset(2).Find(&users)
	fmt.Println(users)
}

// using raw sql
func raw(conn *gorm.DB) {
	var users []User
	conn.Raw("select * from users where ID > ?", 1).Scan(&users)
	fmt.Println(users)
}
