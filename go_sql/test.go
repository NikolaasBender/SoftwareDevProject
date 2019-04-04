package main

import (
  "fmt"
  "github.com/SoftwareDevProject/go_sql/go_dev"
  "database/sql"
  _ "github.com/lib/pq"
)

func main() {

  var db *sql.DB

  db = go_dev.Initialize()

  if(db != nil) {
    fmt.Println("It worked")
  }

  var worked bool

  worked = go_dev.GetUserInfo(, db)

  if(worked != false){
    fmt.Println("User info obtained.")
  }

  db.Close()
}
