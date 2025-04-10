package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// load the env file
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}
}

// construct the postgres connection string
func getConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"unique;not null"`
	Age       int
	CreatedAt int64
}

// perform the migrations
func setup(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func main() {
	now := time.Now()
	loadEnv()

	dns := getConnectionString()
	conn, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	setup(conn)
	fmt.Println("Database created successfully..")

	//create a new record
	// user := User{Name: "brandon", Email: "brandon@gmail.com"}

	//batch insertions
	// user := []User{
	// 	{Name: "john", Email: "john@gmail.com"}, {Name: "doe", Email: "doe@gmail.com"},
	// }
	// conn.Create(&user)

	//read record
	selectById(conn)
	getByEmail(conn)
	getAll(conn)
	getColumn(conn)
	getOrderBy(conn)
	pagination(conn)
	raw(conn)

	//update records
	singleColumn(conn)
	multipleColumn(conn)

	//delete record
	delete(conn)
	fmt.Println("time elapsed..", time.Since(now))
}
