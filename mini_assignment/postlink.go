package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// PostLink for post
type PostLink struct {
	ID            int
	CreationDate  string
	PostID        int
	RelatedPostID int
	LinkTypeID    int
}

func insertPostLink(postlink PostLink) {
	_, err := db.Exec("INSERT INTO post_link(id,creation_date,post_id,related_post_id,link_type_id) VALUES($1,$2,$3,$4,$5);", postlink.ID, postlink.CreationDate, postlink.PostID, postlink.RelatedPostID, postlink.LinkTypeID)
	if err != nil {
		//fmt.Println(err)
	} else {
		//fmt.Println("1 Row Inserted into postlinks")
	}
}

func insertPostLinkFromJSON(path string) {
	file, _ := ioutil.ReadFile(path)
	postlinks := []PostLink{}
	_ = json.Unmarshal([]byte(file), &postlinks)

	for _, postlink := range postlinks {
		insertPostLink(postlink)
	}
	fmt.Println("Done.")

}
