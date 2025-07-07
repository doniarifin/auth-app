package database

import (
	"database/sql"
	"log"
)

func RunMigration(db *sql.DB) {
	query := `
	IF NOT EXISTS (
		SELECT * FROM sysobjects WHERE name='users' AND xtype='U'
	)
	CREATE TABLE users (
		id INT IDENTITY(1,1) PRIMARY KEY,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at DATETIME DEFAULT GETDATE()
	)
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed migrate users tabel: %v", err)
	}
	log.Println("migrate users table finished!")
}
