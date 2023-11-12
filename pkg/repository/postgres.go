package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	productsTable = "products"
)

type Config struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSLMode string
}

func NewPostgresDB(cnf Config) (*sqlx.DB, error) {
	db, error := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	cnf.Host, cnf.Port, cnf.Username, cnf.Password, cnf.DBName, cnf.SSLMode))
	if error!=nil {
		return nil, error
	}
	error = db.Ping()
	if error != nil {
		return nil, error
	}

	return db, nil
}