package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// load the env file
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}
}

func main() {
	loadEnv()
	fmt.Println(os.Getenv("DB_USER"))

}
