package main

import (
	"fmt"
	"log"
	"os"

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

type user struct {
	gorm.Model
	name  string
	email string
}

type channel struct {
	gorm.Model
	name        string
	description string
}

type message struct {
	gorm.Model
	content   string
	userID    uint
	channedID uint
	user      user
	channel   channel
}

// perform the migrations
func setup(db *gorm.DB) {
	db.AutoMigrate(&user{}, &channel{}, &message{})
}

func main() {
	loadEnv()

	dns := getConnectionString()
	conn, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	setup(conn)
	fmt.Println("Database created successfully..")

}
