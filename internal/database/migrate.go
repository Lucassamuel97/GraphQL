package database

import "database/sql"

func Migrate(db *sql.DB) {
	createUserTable := `
    CREATE TABLE IF NOT EXISTS categories (
        id TEXT PRIMARY KEY,
        name TEXT,
		description TEXT
    );`

	_, err := db.Exec(createUserTable)
	if err != nil {
		panic(err)
	}
}
