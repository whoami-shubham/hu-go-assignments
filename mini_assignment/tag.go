package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Tag for post
type Tag struct {
	ID            int
	TagName       string
	Count         int
	ExcerptPostID int
	WikiPostID    int
}

// Question struct
type Question struct {
	ID    int
	Title string
	Body  string
	Tags  string
}

func tagHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	username, password, ok := r.BasicAuth()
	if !ok || !authenticate(username, password) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var tag struct {
		Tags []string `json:"tags"`
	}
	err := decoder.Decode(&tag)
	if err != nil || len(tag.Tags) == 0 {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	patterns := ""
	for index, pattern := range tag.Tags {
		if index != 0 {
			patterns += "AND tags like " + "'%<" + pattern + ">%'"
		} else {
			patterns += "tags like  '%<" + pattern + ">%'"
		}
	}
	query := "select id,title,body,tags from posts where " + patterns
	data := []Question{}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request No question Found with that tag", http.StatusBadRequest)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var question Question
		err := rows.Scan(&question.ID, &question.Title, &question.Body, &question.Tags)
		if err != nil {
			log.Fatal(err)
		} else {
			data = append(data, question)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(getSampleJSONData(data))
}

func insertTag(tag Tag) {
	_, err := db.Exec("INSERT INTO tags(id,tag_name,count,excerpt_post_id,wiki_post_id) VALUES($1,$2,$3,$4,$5);", tag.ID, tag.TagName, tag.Count, tag.ExcerptPostID, tag.WikiPostID)
	if err != nil {
		//fmt.Println(err)
	} else {
		//fmt.Println("1 Row Inserted into tags")
	}
}

func insertTagFromJSON(path string) {
	file, _ := ioutil.ReadFile(path)
	tags := []Tag{}
	_ = json.Unmarshal([]byte(file), &tags)

	for _, tag := range tags {
		insertTag(tag)
	}
	fmt.Println("Done.")

}
