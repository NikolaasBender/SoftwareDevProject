package go_dev

import(
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

func addUser(username,password,email,name string,db *sql.DB) (bool) {

  sqlStatement := `INSERT INTO user_login (username, password, email)
  VALUES ($1, $2, $3);`
  sqlStatement := `INSERT INTO user_info (username, name)
  VALUES ($1, $4);`

  err = *db.QueryRow(sqlStatement,username,password,email)

  if err != nil {
    return false
  }

  err = *db.QueryRow(sqlStatement,username,name,bio)

  if(err != nil) {
    return false
  }

  return true
}

func exists(uername string,db *sql.DB) (bool) {

    sqlStatement := `SELECT username FROM user_login
    WHERE username = $1;`

    var uname string

    err = *db.QueryRow(sqlStatement,username).Scan(&uname)

    if(err == sql.ErrNoRows) {
      return false
    }
    else if (err != nil) {
      return false
    }

    return true
}

func validate(username string, password int, db *sql.DB) (bool) {

  sqlStatement := `SELECT username FROM user_login
  WHERE username = $1 AND password = $2;`

  var uname string

  err = *db.QueryRow(sqlStatement,username, password).Scan(&uname)

  if(err == sql.ErrNoRows) {
    return false
  }
  else if (err != nil) {
    return false
  }

  return true
}

func getUserInfo(username string, db *sql.DB) (bool) {

  sqlStatement := `SELECT * FROM user_info
  WHERE username = $1;`

  var (
    uname string
    name string
    bio string
    profileimg string
    bannerimg string
  )

  err = *db.QueryRow(sqlStatement,username).Scan(&uname,&name,&bio,&profileimg,&bannerimg)

  if(err == sql.ErrNoRows) {
    return false
  }
  else if (err != nil) {
    return false
  }

  return
  //Need to figure out the best way to return some of this information.
}

func editUserInfo(username,field, edit string, db *sql.DB) (bool) {

  sqlStatement := `UPDATE user_info
  SET $1 = $2
  WHERE username = $3;`

  var uname string

  err = *db.QueryRow(sqlStatement,field,edit,username).Scan(&uname)

  if(err == sql.ErrNoRows) {
    return false
  }
  else if (err != nil) {
    return false
  }

  return true
}
