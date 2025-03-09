package main

import (
	"database/sql"
	"fmt"

	"github.com/carlosghabrous/lenslocked/models"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database")
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
      id SERIAL PRIMARY KEY,
      name TEXT,
      email TEXT UNIQUE NOT NULL
      );

      CREATE TABLE IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        amount INT,
        description TEXT
      );
    `)

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully created tables")
	id := 23
	row := db.QueryRow(`
    SELECT name, email FROM users WHERE id = $1
    `, id)
	var name, email string
	err = row.Scan(&name, &email)
	if err == sql.ErrNoRows {
		panic(err)
	}
	fmt.Printf("Name: %s, email: %s\n", name, email)
}
