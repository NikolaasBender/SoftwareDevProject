package go_dev

import(
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
)

func AddUser(username string,password int,email,name string, db *sql.DB) (bool) {

  sqlStatement := `INSERT INTO user_login (username, password, email)
  VALUES ($1, $2, $3);`
  sqlStatement2 := `INSERT INTO user_info (username, name)
  VALUES ($1, $2);`

  var err error

  _, err = db.Exec(sqlStatement,username,password,email)

  if err != nil {
    fmt.Println(err)
    return false
  }

  _, err = db.Exec(sqlStatement2,username,name)

  if(err != nil) {
    fmt.Println(err)
    return false
  }

  return true
}

func Exists(username string,db *sql.DB) (bool) {

    sqlStatement := `SELECT username FROM user_login
    WHERE username = $1;`

    var uname string
    var err error

    err = db.QueryRow(sqlStatement,username).Scan(&uname)

    if(err == sql.ErrNoRows) {
      return false
    } else if (err != nil) {
      return false
    }

    return true
}

func Validate(username string, password int, db *sql.DB) (bool) {

  sqlStatement := `SELECT username FROM user_login
  WHERE username = $1 AND password = $2;`

  var uname string
  var err error

  err = db.QueryRow(sqlStatement,username, password).Scan(&uname)

  if(err == sql.ErrNoRows) {
    return false
  } else if (err != nil) {
    return false
  }

  return true
}

func GetUserInfo(username string, db *sql.DB) (bool) {

  sqlStatement := `SELECT * FROM user_info
  WHERE username = $1;`

  var (
    uname string
    name string
    bio sql.NullString
    profileimg sql.NullString
    bannerimg sql.NullString
    err error
  )

  err = db.QueryRow(sqlStatement,username).Scan(&uname,&name,&bio,&profileimg,&bannerimg)

  if(err == sql.ErrNoRows) {
    fmt.Println(err)
    return false
  } else if (err != nil) {
    fmt.Println(err)
    return false
  }

  return true
  //Need to figure out the best way to return some of this information.
}

func EditUserInfo(username,field, edit string, db *sql.DB) (bool) {

  sqlStatement := `UPDATE user_info
  SET $1 = $2
  WHERE username = $3;`

  var uname string
  var err error

  err = db.QueryRow(sqlStatement,field,edit,username).Scan(&uname)

  if(err == sql.ErrNoRows) {
    return false
  } else if (err != nil) {
    return false
  }

  return true
}
