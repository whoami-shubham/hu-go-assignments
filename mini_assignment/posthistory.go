package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// PostHistory for post
type PostHistory struct {
	ID                int
	PostHistoryTypeID int
	PostID            int
	RevisionGUI       string
	CreationDate      string
	UserID            int
	Text              string
	UserDisplayName   string
	Comment           string
}

func insertPostHistory(posthistory PostHistory) {
	_, err := db.Exec("INSERT INTO post_history(id,post_history_type_id,post_id,revision_guid,creation_date,text,user_id,user_display_name,comment) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9);", posthistory.ID, posthistory.PostHistoryTypeID, posthistory.PostID, posthistory.RevisionGUI, posthistory.CreationDate, posthistory.Text, posthistory.UserID, posthistory.UserDisplayName, posthistory.Comment)
	if err != nil {
		//fmt.Println(err)
	} else {
		//fmt.Println("1 Row Inserted into posthistorys")
	}
}

func insertPostHistoryFromJSON(path string) {
	file, _ := ioutil.ReadFile(path)
	posthistorys := []PostHistory{}
	_ = json.Unmarshal([]byte(file), &posthistorys)

	for _, posthistory := range posthistorys {
		insertPostHistory(posthistory)
	}
	fmt.Println("Done.")

}
