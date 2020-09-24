package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

const tagName = "db"

// GetSampleJSONData ... function to convert struct to json
func GetSampleJSONData(data interface{}) []byte {
	dataJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return []byte("got error while converting data into JSON ")
	}
	return dataJSON
}

// Authenticate ... Basic auth
func Authenticate(username, password string, db *sql.DB) (bool, error) {
	var userName string
	row := db.QueryRow("select email from users where email=$1 and password=$2;", strings.TrimSpace(username), password)
	err := row.Scan(&userName)
	if err != nil || userName == "" {
		log.Println(err)
		return false, err
	}
	return true, nil

}

// InsertIntoDb ... generic function to insert all feilds in db
func InsertIntoDb(data []interface{}, dataType interface{}, tableName string, db *sql.DB) error {

	_, err := db.Exec(InsertQuery(tableName, dataType), data...)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	return nil

}

// InsertQuery ... custome Insert query
func InsertQuery(tableName string, dataType interface{}) string {

	t := reflect.TypeOf(dataType)

	query := "INSERT INTO " + tableName + " ( "
	values := " VALUES ( "
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		if i < t.NumField()-1 {
			query += tag + ", "
			values += "$" + strconv.Itoa(i+1) + " , "
		} else {
			query += tag + " "
			values += "$" + strconv.Itoa(i+1) + " "
		}

	}
	values += " ) "
	query += " )" + values
	return query

}

// GetRole ... get role of user
func GetRole(db *sql.DB, email string) (int, string, error) {
	var role string
	var id int
	row := db.QueryRow("select id,role from users where email=$1;", strings.TrimSpace(email))
	err := row.Scan(&id, &role)
	return id, role, err
}

func insertIssueFromJSON(path string, db *sql.DB) error {
	file, err := ioutil.ReadFile(path)
	issues := []Issue{}
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(file), &issues)
	if err != nil {
		return err
	}
	for _, issue := range issues {
		InsertIntoDb(issue.mapStructToArray(), issue, "issues", db)
	}

	fmt.Println("Done.")
	return nil

}

func (x Issue) mapStructToArray() []interface{} {

	t := reflect.ValueOf(&x).Elem()
	values := make([]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Interface()
		values[i] = field
	}
	return values
}

func insertUserFromJSON(path string, db *sql.DB) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	users := []User{}
	err = json.Unmarshal([]byte(file), &users)
	if err != nil {
		return err
	}
	for _, user := range users {
		InsertIntoDb(user.mapStructToArray(), user, "users", db)
	}
	fmt.Println("Done.")
	return nil

}

func (x User) mapStructToArray() []interface{} {

	t := reflect.ValueOf(&x).Elem()
	values := make([]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Interface()
		values[i] = field
	}
	return values
}

func insertProjectFromJSON(path string, db *sql.DB) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	projects := []Project{}
	err = json.Unmarshal([]byte(file), &projects)
	if err != nil {
		return err
	}
	for _, project := range projects {
		InsertIntoDb(project.mapStructToArray(), project, "projects", db)
	}
	fmt.Println("Done.")
	return nil

}

func (x Project) mapStructToArray() []interface{} {

	t := reflect.ValueOf(&x).Elem()
	values := make([]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Interface()
		values[i] = field
	}
	return values
}

// GetOpenIssues ...
func GetOpenIssues(db *sql.DB) ([]Report, error) {
	var issues []Report
	rows, err := db.Query("select id,title,description,assignee,reporter,created_on,updated_on from issues where status=$1;", "OPEN")
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return issues, err
	}

	for rows.Next() {
		var issue Issue
		var ID int
		var email, name, CreatedOn, UpdatedOn string
		err := rows.Scan(&ID, &issue.Title, &issue.Description, &issue.Assignee, &issue.Reporter, &CreatedOn, &UpdatedOn)
		if err != nil {
			log.Println(err)
			return issues, err
		}
		email, _ = GetEmail(db, issue.Assignee)
		name, _ = GetName(db, issue.Reporter)
		issues = append(issues, Report{Email: email, ID: int32(ID), Title: issue.Title, Description: issue.Description, Reporter: name, CreatedOn: CreatedOn, UpdatedOn: UpdatedOn})

	}
	return issues, nil

}

// GetEmail ...
func GetEmail(db *sql.DB, userID int) (string, error) {
	var email string
	row := db.QueryRow("select email from users where id=$1;", userID)
	err := row.Scan(&email)
	if err != nil {
		//log.Println(err)
		return email, err
	}
	return email, nil
}

// GetName ...
func GetName(db *sql.DB, userID int) (string, error) {
	var name string
	row := db.QueryRow("select name from users where id=$1;", userID)
	err := row.Scan(&name)
	if err != nil {
		//log.Println(err)
		return name, err
	}
	return name, nil
}

// GetWatchers ...
func GetWatchers(db *sql.DB, issueID, assignee, reporter int) ([]string, error) {
	var emails []string
	assigneeEmail, _ := GetEmail(db, assignee)
	reporterEmail, _ := GetEmail(db, reporter)
	emails = append(emails, assigneeEmail)
	emails = append(emails, reporterEmail)
	rows, err := db.Query("select user_id from watchers where issue_id=$1;", issueID)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return emails, err
	}
	for rows.Next() {
		var email string
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return emails, err
		}
		email, _ = GetEmail(db, id)
		emails = append(emails, email)
	}
	return emails, nil
}
