package go_dev

import(
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

func addUser(username,password,email,name) (bool) {

  sqlStatement := `INSERT INTO user_login (username, password, email)
  VALUES ($1, $2, $3);`
  sqlStatement := `INSERT INTO user_info (username, name)
  VALUES ($1, $4);`

  err = db.QueryRow(sqlStatement,username,password,email)

  if err != nil {
    return false
  }

  err = db.QueryRow(sqlStatement,username,name,bio)

  if(err != nil) {
    return false
  }

  return true
}

func exists(uername)

func validate(username, password)

func getUserInfo(username)

func editUserInfo()
