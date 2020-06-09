package conn

import "database/sql"

// Conn ...
type Conn struct {
	Mysql *sql.DB
}

// NewConnection ...
func NewConnection() *Conn {
	return &Conn{
		Mysql: Mysql(),
	}
}
