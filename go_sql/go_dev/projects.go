package go_dev

import (
	"database/sql"
	//"fmt"
	_ "github.com/lib/pq"
)

func CreateProject(owner, name string, db *sql.DB) bool {

	sqlStatement := `INSERT INTO projects(owner,name)
  VALUES ($1, $2)`

	var err error

	_, err = db.Exec(sqlStatement, owner, name)

	if err != nil {
		return false
	}

	return true
}

func AddProjectMembers(owner, name, newuser string, db *sql.DB) bool {

	sqlStatement := `UPDATE projects
  SET users = users || '{$1}'
  WHERE owner = $2 AND name = $3;`

	var err error

	_, err = db.Exec(sqlStatement, newuser, owner, name)

	if err != nil {
		return false
	}

	return true
}

func DeleteProject(owner, name string, db *sql.DB) bool {

	sqlStatement := `DELETE FROM projects
  WHERE owner = $1 AND name = $2;`

	var err error

	_, err = db.Exec(sqlStatement, owner, name)

	if err != nil {
		return false
	}

	return true
}
