package go_dev

import(
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

func addUser(username,password,email,name,db) (bool) {

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

func exists(uername,db) (bool) {

    sqlStatement := `SELECT username FROM user_login
    WHERE username = $1;`

    var uname string

    err = db.QueryRow(sqlStatement,username).Scan(&uname)

    if(err == sql.ErrNoRows) {
      return false
    }
    else if (err != nil) {
      return false
    }

    return true
}

func validate(username, password, db) (bool) {

  sqlStatement := `SELECT username FROM user_login
  WHERE username = $1 AND password = $2;`

  var uname string

  err = db.QueryRow(sqlStatement,username, password).Scan(&uname)

  if(err == sql.ErrNoRows) {
    return false
  }
  else if (err != nil) {
    return false
  }

  return true
}

func getUserInfo(username, db) (bool) {

  sqlStatement := `SELECT * FROM user_info
  WHERE username = $1;`

  var 

}

func editUserInfo()
