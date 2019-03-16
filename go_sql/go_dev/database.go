package go_dev

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
  dbname   = "sdev_db"
)

func initialize() (*sql.DB) {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)

  if err != nil {
    return nil
  }

  err = db.Ping()
  if err != nil {
    return nil
  }

  return &db
}
