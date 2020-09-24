package gql

import (
	"database/sql"
	"fmt"

	"github.com/graphql-go/graphql"
)

var db *sql.DB = nil
var loggedUser CurUser

// GetSchema ...
func GetSchema(dB *sql.DB, id int, username, role string) (*graphql.Schema, error) {
	db = dB
	loggedUser.ID = id
	loggedUser.Email = username
	loggedUser.Role = role
	rootQuery := NewRootQuery()
	rootMutation := NewRootMutation()
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQuery.Query,
			Mutation: rootMutation.Mutation,
		},
	)
	if err != nil {
		fmt.Println("error in GetSchema ", err)
		return nil, err
	}
	return &schema, nil
}
