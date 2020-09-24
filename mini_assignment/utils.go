package main

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
)

// SORTBY parameters
var SORTBY = []string{"score", "creation_date", "last_activity_date"}

const tagName = "db"

var isPresent = func(sortby string) bool {
	for _, value := range SORTBY {
		if value == sortby {
			return true
		}
	}
	return false
}

// getSampleJSONData ... function to convert struct to json
func getSampleJSONData(data interface{}) []byte {
	dataJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return []byte("got error while converting data into JSON ")
	}
	return dataJSON
}

func authenticate(username, password string) bool {
	var userName string
	row := db.QueryRow("select username from customer where username=$1 and password=$2;", username, password)
	err := row.Scan(&userName)
	if err != nil || userName == "" {
		log.Println(err)
		return false
	}
	return true

}

func insert(tableName string, x interface{}) string {

	t := reflect.TypeOf(x)

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
