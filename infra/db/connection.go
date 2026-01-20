package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	// user -> postgres
	// password -> alvinsakib
	// host -> localhost
	// port -> 5432
	// db name ->
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
	if !cnf.EnableSSLMODE {
		connString += " sslmode=disable"
	}
	return connString
	// return "user=postgres password=alvinsakib host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource) // dbCon => client
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
