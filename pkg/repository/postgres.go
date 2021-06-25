package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

const usersTable = "users"

func NewPostgresDB() (*sqlx.DB, error) {
	dbport, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		logrus.Fatalf("Unable to convert the string into int.  %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), dbport, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	logrus.Println("Successfully connected to the database!")

	return db, nil

}
