package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	// Set file permissions (read/write for owner only)
	os.OpenFile("mydb.sqlite", os.O_RDWR|os.O_CREATE, 0600)

	// Open the database
	key := getSecureKey() // Implement this function to securely retrieve or generate your key
	connectionString := fmt.Sprintf("file:mydb.sqlite?_pragma_key=%s&_pragma_busy_timeout=5000&_pragma_journal_mode=WAL&_pragma_synchronous=NORMAL&_pragma_foreign_keys=ON&_pragma_temp_store=MEMORY&cache=shared&mode=rwc", key)

	db, err := sql.Open("sqlite", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// It's a good practice to set a connection limit
	db.SetMaxOpenConns(1) // SQLite supports only one writer at a time

	// Create a table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)`)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a record
	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", "John Doe")
	if err != nil {
		log.Fatal(err)
	}

	// Query the database
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User: %d, %s\n", id, name)
	}
}

func getSecureKey() string {
	// Implement secure key retrieval
	return "your-very-secure-and-long-key"
}
