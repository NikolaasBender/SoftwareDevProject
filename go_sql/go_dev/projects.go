package go_dev

import(
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

func createProject(owner, name string, db) (bool) {

  sqlStatement := `INSERT INTO projects(owner,name)
  VALUES ($1, $2)`

  err = db.QueryRow(sqlStatement,owner,name)

  if(err != nil) {
    return false
  }

  return true
}

func addMembers(owner, name, newuser string, db) (bool) {

  sqlStatement := `UPDATE projects
  SET users = users || '{$1}'
  WHERE owner = $2 AND name = $3;`

  err = db.QueryRow(sqlStatement,newuser,owner,name)

  if(err != nil)
  {
    return false
  }

  return true
}

func deleteProject(owner, name string, db) (bool) {

  sqlStatement := `DELETE FROM projects
  WHERE owner = $1 AND name = $2;`

  err = db.QueryRow(sqlStatement, owner, name)

  if(err != nil)
  {
    return false
  }

  return true
}
