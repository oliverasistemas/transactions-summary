package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	envMySQLUser     = "MYSQL_USER"
	envMySQLPassword = "MYSQL_PASSWORD"
	envMySQLHost     = "MYSQL_HOST"
	envMySQLPort     = "MYSQL_PORT"
	envMySQLDatabase = "MYSQL_DB"
)

// ConnectMySQL establishes a connection to the MySQL database.
func ConnectMySQL() (*sql.DB, error) {
	connectionString := getConnectionString()

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func getConnectionString() string {
	user := os.Getenv(envMySQLUser)
	password := os.Getenv(envMySQLPassword)
	host := os.Getenv(envMySQLHost)
	port := os.Getenv(envMySQLPort)
	database := os.Getenv(envMySQLDatabase)

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
}
