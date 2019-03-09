package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "password"
  dbname   = "test_db"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
  
  sqlStatement := `CREATE TABLE IF NOT EXISTS users (
  	id SERIAL PRIMARY KEY,
  	username TEXT UNIQUE NOT NULL,
  	password TEXT NOT NULL,
	first_name TEXT,
	last_name TEXT,
  	email TEXT
	);` 
  _, err = db.Exec(sqlStatement)
  if err != nil {
	  panic(err)
  }
  sqlStatement2 := `
        INSERT INTO users (username, password, first_name, last_name, email)
        VALUES ('user1', 'topSecret','Jane','Smith','jane@example.com') ON CONFLICT DO NOTHING;`
  _, err = db.Exec(sqlStatement2)
  if err != nil {
    panic(err)
  }



}

