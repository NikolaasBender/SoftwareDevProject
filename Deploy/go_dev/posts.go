package go_dev

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
)
func createPost(parentTask_name, title, user, db *sql.DB) (bool){
	sqlStatement1:= `SELECT id FROM tasks WHERE name = $1;`
	
	
	var parentID string
  	var err error
	
	err = db.QueryRow(sqlStatement1,parentTask_name).Scan(&parentID)

 	if(err == sql.ErrNoRows) {
   		return false
  	} else if (err != nil) {
    	return false
  	}

 	sqlStatement2 := `INSERT INTO posts(task,title,users)
  	VALUES ($1, $2, $3)`

  	_, err = db.Exec(sqlStatement2,parentID, title, user)

  	if(err != nil) {
    	return false
  	}

  	return true
}
func addContentPost(title, user, content, db *sql.DB)(bool){
	sqlStatement:= `UPDATE posts
  	SET content = $1
  	WHERE title = $2 AND user = $3;`
	
	_, err = db.Exec(sqlStatement,content, title, user)

  	if(err != nil) {
    	return false
  	}

  	return true
}

func deletePost(title, user, db *sql.DB) (bool) {
  	sqlStatement := `DELETE FROM posts
  	WHERE title = $1 AND users = $2;`

  	var err error

  	_, err = db.Exec(sqlStatement, title, user)

  	if(err != nil) {
    	return false
  	}

  	return true
}
