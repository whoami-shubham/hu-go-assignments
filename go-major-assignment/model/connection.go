package model

import (
	"database/sql"
	"fmt"
)

// ...
const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "shubham"
	PASSWORD = "password"
	DBNAME   = "issuetracker"
)

// ConnStr ...
var ConnStr = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s",
	HOST, PORT, USER, PASSWORD, DBNAME)

// for test db
var testConnStr = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s",
	HOST, PORT, USER, PASSWORD, "test_"+DBNAME)

// InsertJSON ...
func InsertJSON(db *sql.DB) {
	insertUserFromJSON("./model/json/users.json", db)
	insertProjectFromJSON("./model/json/projects.json", db)
	insertIssueFromJSON("./model/json/issues.json", db)
}
