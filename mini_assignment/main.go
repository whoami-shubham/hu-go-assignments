package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB = nil

func init() {
	var err error
	db, err = sql.Open("postgres",
		connStr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}
}

func main() {
	port := ":4000"
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	insertJSON()
	http.HandleFunc("/posts", postHandler)       // GET /posts to get all posts
	http.HandleFunc("/comments", commentHandler) // GET /comments?id=post_id to get all comments of a post
	http.HandleFunc("/upvote", upVoteHandler)
	http.HandleFunc("/downvote", downVoteHandler)
	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/customer", customerHandler)
	http.HandleFunc("/tags", tagHandler)
	log.Fatal(http.ListenAndServe(port, nil))

	defer db.Close()

}
