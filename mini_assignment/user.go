package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// User for post
type User struct {
	ID              int
	Reputation      int
	CreationDate    string
	DisplayName     string
	LastAccessDate  string
	WebsiteURL      string
	Location        string
	AboutMe         string
	Views           int
	UpVotes         int
	DownVotes       int
	AccountID       int
	ProfileImageURL string
}

// UserHandler ...
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	username, password, ok := r.BasicAuth()
	if !ok || !authenticate(username, password) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}
	id := r.URL.Query().Get("id")
	data := []User{}
	if id != "" {
		var user User
		rows, err := db.Query("select * from users where id = $1", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&user.ID, &user.Reputation, &user.CreationDate, &user.DisplayName, &user.LastAccessDate, &user.WebsiteURL, &user.Location, &user.AboutMe, &user.Views, &user.UpVotes, &user.DownVotes, &user.AccountID, &user.ProfileImageURL)
			if err != nil {
				log.Fatal(err)
			}
		}
		data = append(data, user)
	} else {
		rows, err := db.Query("select * from users;")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Reputation, &user.CreationDate, &user.DisplayName, &user.LastAccessDate, &user.WebsiteURL, &user.Location, &user.AboutMe, &user.Views, &user.UpVotes, &user.DownVotes, &user.AccountID, &user.ProfileImageURL)
			if err != nil {
				log.Fatal(err)
			} else {
				data = append(data, user)
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(getSampleJSONData(data))

}

func insertUser(user User) {
	// _, err := db.Exec(insert("users", user), user)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	//fmt.Println("1 Row Inserted into users")
	// }
}

func insertUserFromJSON(path string) {
	file, _ := ioutil.ReadFile(path)
	users := []User{}
	_ = json.Unmarshal([]byte(file), &users)

	for _, user := range users {
		insertUser(user)
	}

}
