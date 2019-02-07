package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type (
	// App ...
	App struct {
		Connection *sql.DB
	}
)

func main() {
	app := App{}
	app.Connection = Mysql()
}

// Mysql ...
func Mysql() *sql.DB {
	conn, err := sql.Open("mysql", "palmirinha:palmirinha@tcp(palmirinha-db:3306)/palmirinha")
	if err != nil {
		fmt.Println("Conn Mysql: ", err)
	}

	return conn
}
