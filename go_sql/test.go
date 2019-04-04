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
    fmt.Println("It worked.\n")
  }

  var worked bool

  worked = go_dev.AddUser("Ben67", 2901, "ben.singer@colorado.edu", "Ben", db)

  if(worked != false) {
    fmt.Println("User added.\n")
  }else{
    fmt.Println("Failed to add.")
  }

  worked = go_dev.GetUserInfo("RS-Coop", db)

  if(worked != false){
    fmt.Println("User info obtained.\n")
  }else{
    fmt.Println("Failed to obtain.")
  }

  db.Close()
}
