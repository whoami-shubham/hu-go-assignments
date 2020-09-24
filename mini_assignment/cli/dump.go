package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"io/ioutil"

	_ "github.com/lib/pq"
)

type Question struct {
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

const (
	host   = "localhost"
	port   = 5432
	dbname = "test_overflow"
)

func main() {
	var username, password string
	curTime := time.Now()
	format := "15:04:05.99999999"
	fileName := "unanswered_ques.json" + curTime.Format(format)

	fmt.Print("username : ")
	_, err := fmt.Scan(&username)
	if err != nil {
		fmt.Println(err)
	}
	print("\033[H\033[2J")
	fmt.Print("password : ")
	_, err = fmt.Scan(&password)
	if err != nil {
		fmt.Println(err)
	}
	print("\033[H\033[2J")
	if username != "" && password != "" {
		var connStr = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s",
			host, port, username, password, dbname)
		db, err := sql.Open("postgres",
			connStr)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Connecting to database....")
		}
		print("\033[H\033[2J")
		query := "select * from posts where parent_id=0 and answer_count=0"
		fmt.Println("Fetching result from Database ...")
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()
		data := []Question{}
		for rows.Next() {
			var post Question
			err := rows.Scan(&post.ID, &post.PostTypeID, &post.Score, &post.ViewCount, &post.Tags, &post.AnswerCount, &post.CommentCount, &post.FavoriteCount, &post.CreationDate, &post.Body, &post.ClosedDate, &post.AcceptedAnswerID, &post.ParentID, &post.OwnerUserID, &post.OwnerDisplayName, &post.LastEditorUserID, &post.LastEditorDisplayName, &post.LastEditDate, &post.LastActivityDate, &post.Title, &post.CommunityOwnedDate)
			if err != nil {
				log.Println(err)
				return
			} else {
				data = append(data, post)
			}
		}
		createJSONFile(data, fileName)

	} else {
		fmt.Println("invalid username or password.")
	}

}

func createJSONFile(data []Question, fileName string) {
	dataJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("got error while converting data into JSON ")
		return
	}
	err = ioutil.WriteFile(fileName, dataJSON, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	print("\033[H\033[2J")
	fmt.Println("Done.")
}
