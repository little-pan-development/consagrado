package models

import "database/sql"

type (
	// App ...
	App struct {
		Connection *sql.DB
	}
)
