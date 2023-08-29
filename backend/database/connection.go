package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {
	log.Println("Initializing database...")

	db, err := sql.Open("sqlite", "./notesapp.db")
	if err != nil {
		log.Println("Error opening database:", err)
		return nil, err
	}

	// Проверка наличия таблиц
	if !tablesExist(db) {
		log.Println("Starting transaction...")
		tx, err := db.Begin()
		if err != nil {
			log.Println("Error starting transaction:", err)
			return nil, err
		}

		createUsersTable := `
            CREATE TABLE IF NOT EXISTS users (
                id INTEGER PRIMARY KEY,
                username TEXT UNIQUE,
                password TEXT
            );
        `
		log.Println("Creating users table...")
		_, err = tx.Exec(createUsersTable)
		if err != nil {
			log.Println("Error creating users table:", err)
			tx.Rollback()
			return nil, err
		}

		createNotesTable := `
            CREATE TABLE IF NOT EXISTS notes (
                id INTEGER PRIMARY KEY,
                content TEXT,
                user_id INTEGER,
                FOREIGN KEY (user_id) REFERENCES users(id)
            );
        `
		log.Println("Creating notes table...")
		_, err = tx.Exec(createNotesTable)
		if err != nil {
			log.Println("Error creating notes table:", err)
			tx.Rollback()
			return nil, err
		}

		log.Println("Committing transaction...")
		err = tx.Commit()
		if err != nil {
			log.Println("Error committing transaction:", err)
			return nil, err
		}
	} else {
		log.Println("Tables already exist, skipping initialization.")
	}

	log.Println("Database initialization completed.")
	return db, nil
}

func tablesExist(db *sql.DB) bool {
	// Проверяем наличие таблиц в базе данных
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		log.Println("Error checking tables existence:", err)
		return false
	}
	defer rows.Close()

	// Проверяем, есть ли нужные таблицы
	tables := []string{"users", "notes"}
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Println("Error scanning table name:", err)
			return false
		}
		for _, t := range tables {
			if tableName == t {
				tables = removeFromSlice(tables, t)
			}
		}
	}

	return len(tables) == 0
}

func removeFromSlice(slice []string, item string) []string {
	for i, s := range slice {
		if s == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
