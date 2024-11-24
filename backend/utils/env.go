package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	DB_HOST = myEnv["DB_HOST"]
	DB_PORT = myEnv["DB_PORT"]
	DB_USER = myEnv["DB_USER"]
	DB_PASS = myEnv["DB_PASS"]

	JWT_SECRET = myEnv["JWT_SECRET"]
}

var (
	DB_HOST = ""
	DB_PORT = ""
	DB_USER = ""
	DB_PASS = ""

	JWT_SECRET = ""
)
