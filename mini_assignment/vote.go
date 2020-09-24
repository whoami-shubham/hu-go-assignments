package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

// Vote for post
type Vote struct {
	ID           int    `db:"id"`
	PostID       int    `db:"post_id"`
	VoteTypeID   int    `db:"vote_type_id"`
	CreationDate string `db:"creation_date"`
	UserID       int    `db:"user_id"`
	BountyAmount int    `db:"bounty_amount"`
}

func upVoteHandler(w http.ResponseWriter, r *http.Request) {
	updateVotes("upvotes", 1, w, r)
}

func downVoteHandler(w http.ResponseWriter, r *http.Request) {
	updateVotes("downvotes", -1, w, r)
}

func insertVote(vote []interface{}, x interface{}) {
	//fmt.Println(vote)

	_, err := db.Exec(insert("votes", x), vote...)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println("1 Row Inserted into votes")
	}
}

func insertVoteFromJSON(path string) {
	file, _ := ioutil.ReadFile(path)
	votes := []Vote{}
	_ = json.Unmarshal([]byte(file), &votes)

	for _, vote := range votes {
		insertVote(mapStructToArray(vote), vote)
	}
	fmt.Println("Done.")

}

func updateVotes(operation string, increment int, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("post_id")
	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	username, password, ok := r.BasicAuth()
	if !ok || !authenticate(username, password) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}
	var score, ownerUserID int
	rows, err := db.Query("select score,owner_user_id from posts where id = $1", id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&score, &ownerUserID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	}
	_, err = db.Exec("UPDATE posts set score = $1 where id = $2;", score+increment, id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	var reputation, upvotes int
	query := "select reputation," + operation + " from users where id = $1"
	rows, err = db.Query(query, ownerUserID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&reputation, &upvotes)
		if err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	}
	query = "UPDATE users set reputation = $1," + operation + " = $2 where id = $3;"
	_, err = db.Exec(query, reputation+increment, upvotes+increment, ownerUserID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//fmt.Println("Voted Sucessfully")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))

}

func mapStructToArray(x Vote) []interface{} {

	t := reflect.ValueOf(&x).Elem()
	values := make([]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Interface()
		values[i] = field
	}
	return values
}
