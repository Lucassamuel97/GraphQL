package database

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {
	createCategoriesTable := `
    CREATE TABLE IF NOT EXISTS categories (
        id TEXT PRIMARY KEY,
        name TEXT,
		description TEXT
    );`

	_, err := db.Exec(createCategoriesTable)
	if err != nil {
		log.Fatalf("Error creating categories table: %v", err)
	}

	createCoursesTable := `
    CREATE TABLE IF NOT EXISTS courses (
        id TEXT PRIMARY KEY,
        name TEXT,
		description TEXT,
		category_id TEXT
    );`

	_, err = db.Exec(createCoursesTable)
	if err != nil {
		log.Fatalf("Error creating courses table: %v", err)
	}
}
