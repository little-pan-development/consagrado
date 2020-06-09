package conn

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Mysql ...
func Mysql() *sql.DB {

	database := os.Getenv("MYSQL_DATABASE")
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	port := os.Getenv("MYSQL_PORT")
	host := os.Getenv("MYSQL_HOST")

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci", username, password, host, port, database)

	conn, err := sql.Open("mysql", uri)
	if err != nil {
		fmt.Println("Conn Mysql: ", err)
	}

	return conn
}
