package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Comment for post
type Comment struct {
	ID              int
	PostID          int
	Score           int
	Text            string
	UserID          int
	CreationDate    string
	UserDisplayName string
}

// commentHandler ...
func commentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("post_id")
	if r.Method != http.MethodGet || id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	username, password, ok := r.BasicAuth()
	if !ok || !authenticate(username, password) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}
	data := []Comment{}
	rows, err := db.Query("select * from comments where post_id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.CreationDate, &comment.Text, &comment.UserID, &comment.Score, &comment.UserDisplayName)
		if err != nil {
			log.Fatal(err)
		} else {
			data = append(data, comment)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(getSampleJSONData(data))
}

func insertComment(comment Comment) {
	_, err := db.Exec("INSERT INTO comments(id,post_id,score,text,user_id,creation_date, user_display_name) VALUES($1,$2,$3,$4,$5,$6,$7);", comment.ID, comment.PostID, comment.Score, comment.Text, comment.UserID, comment.CreationDate, comment.UserDisplayName)
	if err != nil {
		//fmt.Println(err)
	} else {
		//fmt.Println("1 Row Inserted into comments")
	}
}

func insertCommentFromJSON(path string) {
	file, _ := ioutil.ReadFile(path)
	comments := []Comment{}
	_ = json.Unmarshal([]byte(file), &comments)

	for _, comment := range comments {
		insertComment(comment)
	}
	fmt.Println("Done.")

}
