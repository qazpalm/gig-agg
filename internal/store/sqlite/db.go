package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func InitialiseSchema(db *sql.DB) error {
	log.Println("Initialising schema...")
	_, err := db.Exec(enableForeignKeys)
	if err != nil {
		return fmt.Errorf("error enabling foreign keys: %w", err)
	}

	// Execute all the schema creation statements
	for _, stmt := range schemaStatements {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Printf("error executing schema statement: %v\nSQL: %s\n", err, stmt)
			return fmt.Errorf("error executing schema statement: %w", err)
		}
	}

	log.Println("Schema initialised successfully")
	return nil
}

