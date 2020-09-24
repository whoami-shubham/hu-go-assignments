package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Post for post
type Post struct {
	ID                    int
	ViewCount             int
	AnswerCount           int
	CommentCount          int
	FavoriteCount         int
	ClosedDate            string
	Title                 string
	PostTypeID            int
	CreationDate          string
	Score                 int
	Body                  string
	Tags                  string
	AcceptedAnswerID      int
	ParentID              int
	OwnerUserID           int
	OwnerDisplayName      string
	LastEditorUserID      int
	LastEditorDisplayName string
	LastEditDate          string
	LastActivityDate      string
	CommunityOwnedDate    string
}

// PostSmall required for display limited data of posts
type PostSmall struct {
	ID            int    `json:"id"`
	ViewCount     int    `json:"view_count"`
	AnswerCount   int    `json:"answer_count"`
	CommentCount  int    `json:"comment_count"`
	FavoriteCount int    `json:"favourite_count"`
	ClosedDate    string `json:"closed_date"`
	Title         string `json:"title"`
}

// Answer type for post containing answers
type Answer struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Body             string `json:"body"`
	Score            int    `json:"score"`
	LastActivityDate string `json:"last_activity_date"`
	CreationDate     string `json:"creation_date"`
}

// PostHandler ...
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	username, password, ok := r.BasicAuth()
	if !ok || !authenticate(username, password) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}
	id := r.URL.Query().Get("post_id")
	sortBy := r.URL.Query().Get("sort_by")
	if sortBy == "" {
		sortBy = "score"
	}
	if id == "" { // if id is not provided return all posts
		data := []PostSmall{}
		rows, err := db.Query("select id,view_count,answer_count,comment_count,favourite_count,closed_date,title from posts;")
		if err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var post PostSmall
			err := rows.Scan(&post.ID, &post.ViewCount, &post.AnswerCount, &post.CommentCount, &post.FavoriteCount, &post.ClosedDate, &post.Title)
			if err != nil {
				log.Fatal(err)
			} else {
				data = append(data, post)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(getSampleJSONData(data))
	} else { // if id is provided return answers of that post
		data := []Answer{}
		query := "select id,title,body,score,last_activity_date,creation_date from posts where parent_id = $1 and parent_id!=0 order by "
		if isPresent(sortBy) {
			query = query + sortBy + " desc;"
		}
		rows, err := db.Query(query, id)
		if err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var post Answer
			err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.Score, &post.LastActivityDate, &post.CreationDate)
			if err != nil {
				log.Println(err)
			} else {
				data = append(data, post)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(getSampleJSONData(data))

	}

}

func insertPost(post Post) {
	_, err := db.Exec("INSERT INTO posts(id,post_type_id,score,view_count,tags,answer_count,comment_count,favourite_count,creation_date,body,closed_date,accepted_answer_id,parent_id,owner_user_id,owner_display_name,last_editor_user_id,last_editor_display_name,last_edit_date,last_activity_date,title,community_owned_date) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21);", post.ID, post.PostTypeID, post.Score, post.ViewCount, post.Tags, post.AnswerCount, post.CommentCount, post.FavoriteCount, post.CreationDate, post.Body, post.ClosedDate, post.AcceptedAnswerID, post.ParentID, post.OwnerUserID, post.OwnerDisplayName, post.LastEditorUserID, post.LastEditorDisplayName, post.LastEditDate, post.LastActivityDate, post.Title, post.CommunityOwnedDate)
	if err != nil {
		//fmt.Println(err)
	} else {
		//fmt.Println("1 Row Inserted into posts")
	}
}

func insertPostFromJSON(path string) {
	file, _ := ioutil.ReadFile(path)
	posts := []Post{}
	_ = json.Unmarshal([]byte(file), &posts)
	for _, post := range posts {
		insertPost(post)
	}
	fmt.Println("Done.")

}
