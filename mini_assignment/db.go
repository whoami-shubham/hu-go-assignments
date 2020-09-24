package main

import (
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "shubham"
	password = "password"
	dbname   = "overflow"
)

var connStr = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s",
	host, port, user, password, dbname)

// insertJSON ...

func insertJSON() {
	//insertUserFromJSON("./json/Users.json")
	// insertPostFromJSON("./json/Posts.json")
	insertVoteFromJSON("./json/Votes.json")
	// insertCommentFromJSON("./json/Comments.json")
	// insertBadgeFromJSON("./json/Badges.json")
	// insertTagFromJSON("./json/Tags.json")
	// insertPostLinkFromJSON("./json/PostLinks.json")
	// insertPostHistoryFromJSON("./json/PostHistory.json")

}
