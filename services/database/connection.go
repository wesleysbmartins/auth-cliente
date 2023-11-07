package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type databaseCredentials struct {
	user          string
	password      string
	database_name string
	port          string
	host          string
	ssl_mode      string
}

func Connection() *sql.DB {

	var credentials databaseCredentials

	credentials.user = os.Getenv("POSTGRES_USER")
	credentials.password = os.Getenv("POSTGRES_PASSWORD")
	credentials.database_name = os.Getenv("POSTGRES_DB")
	credentials.port = os.Getenv("POSTGRES_PORT")
	credentials.host = os.Getenv("POSTGRES_HOST")
	credentials.ssl_mode = "disable"

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		credentials.user,
		credentials.password,
		credentials.database_name,
		credentials.ssl_mode,
		credentials.host,
		credentials.port,
	)

	fmt.Println("CONNSTR: ", connStr)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("ERROR TO CONNECT DATABASE")
		fmt.Println(err)

	}

	return db
}
