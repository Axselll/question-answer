package main

import (
	"fmt"
	"os"
	"qa/config"
	"qa/exception"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	exception.CheckError(err)

	// test only idk
	database := config.NewDB(os.Getenv("ENGINE"), os.Getenv("ADDRESS_DB"))
	s := os.Getenv("ENGINE")
	t := database.Ping()
	fmt.Print(s, t)

}
