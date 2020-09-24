package main

import (
	"database/sql"
	"encoding/json"
	"go-major-assignment/client"
	"go-major-assignment/gql"
	"go-major-assignment/model"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

var db *sql.DB = nil

func init() {
	var err error
	db, err = sql.Open("postgres",
		model.ConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}
}

func main() {
	port := ":4000"
	log.SetFlags(log.LstdFlags | log.Lshortfile) // for comment line number in file in logging
	//model.InsertJSON(db)    // function for insert data from json file to db
	go client.Cron(db)
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		verified, _ := model.Authenticate(username, password, db)
		if !verified {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		id, role, errr := model.GetRole(db, username)
		if errr != nil {
			log.Println(errr)
			//return
		}

		schema, err := gql.GetSchema(db, id, username, role)
		if err != nil {
			log.Println(err)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:        *schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})

	log.Fatal(http.ListenAndServe(port, nil))
	defer db.Close()

}
