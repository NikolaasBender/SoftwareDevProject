package go_dev

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
)

func createTask(project_name,project_owner, task_name string,db *sql.DB)(bool) {
  	sqlStatement1:= `SELECT id FROM projects WHERE owner = $1 AND name = $2;`
	
	
	var parentID string
  	var err error
	
	err = db.QueryRow(sqlStatement1,project_owner,project_name).Scan(&parentID)

 	if(err == sql.ErrNoRows) {
   		return false
  	} else if (err != nil) {
    	return false
  	}

 	sqlStatement2 := `INSERT INTO tasks(project,name,status)
  	VALUES ($1, $2, 0)`

  	_, err = db.Exec(sqlStatement2,parentID,task_name)

  	if(err != nil) {
    	return false
  	}

  	return true
}

func addTaskMembers(project_name,project_owner, task_name,newMember,db *sql.DB)(bool) {
	sqlStatement1:= `SELECT id FROM projects WHERE owner = $1 AND name = $2;`
	
	var parentID string
  	var err error
	err = db.QueryRow(sqlStatement1,project_owner,project_name).Scan(&parentID)

 	if(err == sql.ErrNoRows) {
   		return false
  	} else if (err != nil) {
    	return false
  	}
	
	sqlStatement:= `UPDATE tasks
  	SET users = users || '{$1}'
  	WHERE project = $2 AND name = $3;`

  	_, err = db.Exec(sqlStatement,newMember,parentID,task_name)

  	if(err != nil) {
    	return false
  	}

  	return true
}

func updateStatus(project_name,project_owner, task_name,status int,db *sql.DB)(bool) {
  	sqlStatement1:= `SELECT id FROM projects WHERE owner = $1 AND name = $2;`
	
	var parentID string
  	var err error
	err = db.QueryRow(sqlStatement1,project_owner,project_name).Scan(&parentID)

 	if(err == sql.ErrNoRows) {
   		return false
  	} else if (err != nil) {
    	return false
  	}
	
	sqlStatement:= `UPDATE tasks
  	SET status = $1
  	WHERE project = $2 AND name = $3;`

  	_, err = db.Exec(sqlStatement,status,parentID,task_name)

  	if(err != nil) {
    	return false
  	}

  	return true
}

func addDescription(project_name,project_owner, task_name, description,db *sql.DB) (bool){
  	sqlStatement1:= `SELECT id FROM projects WHERE owner = $1 AND name = $2;`
	
	var parentID string
  	var err error
	err = db.QueryRow(sqlStatement1,project_owner,project_name).Scan(&parentID)

 	if(err == sql.ErrNoRows) {
   		return false
  	} else if (err != nil) {
    	return false
  	}
	
	sqlStatement:= `UPDATE tasks
  	SET description = $1
  	WHERE project = $2 AND name = $3;`

  	_, err = db.Exec(sqlStatement,description,parentID,task_name)

  	if(err != nil) {
    	return false
  	}

  	return true
}

func dueDate(project_name,project_owner, task_name,dueDate,db *sql.DB)(bool) {
  	sqlStatement1:= `SELECT id FROM projects WHERE owner = $1 AND name = $2;`
	
	var parentID string
  	var err error
	err = db.QueryRow(sqlStatement1,project_owner,project_name).Scan(&parentID)

 	if(err == sql.ErrNoRows) {
   		return false
  	} else if (err != nil) {
    	return false
  	}
	
	sqlStatement:= `UPDATE tasks
  	SET due_date = $1
  	WHERE project = $2 AND name = $3;`

  	_, err = db.Exec(sqlStatement,dueDate,parentID,task_name)

  	if(err != nil) {
    	return false
  	}

  	return true
}

func deleteTask(project_name,project_owner, task_name,db *sql.DB) {
  	sqlStatement1:= `SELECT id FROM projects WHERE owner = $1 AND name = $2;`
	
	var parentID string
  	var err error
	err = db.QueryRow(sqlStatement1,project_owner,project_name).Scan(&parentID)

 	if(err == sql.ErrNoRows) {
   		return false
  	} else if (err != nil) {
    	return false
  	}
	
	sqlStatement := `DELETE FROM tasks
  	WHERE  project= $1 AND name = $2;`

  	var err error

  	_, err = db.Exec(sqlStatement, parentID, task_name)

  	if(err != nil) {
    	return false
  	}

  	return true
}
