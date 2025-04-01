package db

import "database/sql"

// Open DB connection
func connect() (*sql.DB, error) {
	connectionString := "xxxxxxxxxxxxxxx"

	db, erro := sql.Open("mydb", connectionString)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
