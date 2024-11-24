package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	SERVER_PORT = myEnv["SERVER_PORT"]

	DB_HOST = myEnv["DB_HOST"]
	DB_PORT = myEnv["DB_PORT"]
	DB_USER = myEnv["DB_USER"]
	DB_PASS = myEnv["DB_PASS"]
	DB_DATABASE = myEnv["DB_DATABASE"]

	JWT_SECRET = myEnv["JWT_SECRET"]
}

var (
	SERVER_PORT string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASS     string
	DB_DATABASE string

	JWT_SECRET string
)
