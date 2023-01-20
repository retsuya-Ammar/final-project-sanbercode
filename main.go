package main

import (
	"database/sql"
	"final-project-sanbercode/database"
	"final-project-sanbercode/routers"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	} else {
		fmt.Println("Successfully loaded .env file")
	}

	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Failed to connect to database!")
		panic(err)
	} else {
		fmt.Println("Successfully connected!")
	}

	database.DBMigrate(DB)

	defer DB.Close()

	routers.SetupRouter().Run(":" + os.Getenv("PORT"))
}
