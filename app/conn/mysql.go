package conn

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Mysql ...
func Mysql() *sql.DB {
	conn, err := sql.Open("mysql", "palmirinha:palmirinha@tcp(palmirinha-db:3306)/palmirinha?charset=utf8mb4&collation=utf8mb4_unicode_ci")
	if err != nil {
		fmt.Println("Conn Mysql: ", err)
	}

	return conn
}
