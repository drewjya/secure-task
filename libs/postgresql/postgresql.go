package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Connection ...

// Connect ...
func Connect() (*sql.DB, error) {
	godotenv.Load(".env")
	DB_HOST := GoDotEnvVariable("DB_HOST")
	DB_USER := GoDotEnvVariable("DB_USER")
	DB_PASSWORD := GoDotEnvVariable("DB_PASSWORD")
	DB_NAME := GoDotEnvVariable("DB_NAME")
	DB_PORT := GoDotEnvVariable("DB_PORT")
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)

	return db, err
}

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
