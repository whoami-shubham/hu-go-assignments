package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Badge for post
type Badge struct {
	ID       int
	UserID   int
	Date     string
	Name     string
	Class    int
	TagBased string
}

func insertBadge(badge Badge) {
	_, err := db.Exec("INSERT INTO badges(id,user_id,date,name, class,tagbased) VALUES($1,$2,$3,$4,$5,$6);", badge.ID, badge.UserID, badge.Date, badge.Name, badge.Class, badge.TagBased)
	if err != nil {
		//fmt.Println(err)
	} else {
		//fmt.Println("1 Row Inserted into badges")
	}
}

func insertBadgeFromJSON(path string) {
	file, _ := ioutil.ReadFile(path)
	badges := []Badge{}
	_ = json.Unmarshal([]byte(file), &badges)

	for _, badge := range badges {
		insertBadge(badge)
	}
	fmt.Println("Done.")

}
