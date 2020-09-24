package model

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

var expected = []string{"INSERT INTO issues ( title, description, type, assignee, reporter, status, project  ) VALUES ( $1 , $2 , $3 , $4 , $5 , $6 , $7  ) ", "INSERT INTO users ( name, password, email, role  ) VALUES ( $1 , $2 , $3 , $4  ) ", "INSERT INTO projects ( name, created_by  ) VALUES ( $1 , $2  ) "}
var output = []string{}

func initDB() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err := sql.Open("postgres",
		testConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		_, err = db.Exec("DELETE FROM projects;")
		if err != nil {
			log.Println(err)
		}
		_, err = db.Exec("DELETE FROM issues;")
		if err != nil {
			log.Println(err)
		}
		_, err = db.Exec("DELETE FROM users;")
		if err != nil {
			log.Println(err)
		}
	}
}
func TestInsertQuery(t *testing.T) {
	output = append(output, InsertQuery("issues", Issue{}))
	output = append(output, InsertQuery("users", User{}))
	output = append(output, InsertQuery("projects", Project{}))
	if expected[0] != output[0] {
		log.Println("expected : ", expected[0])
		log.Println("got : ", output[0])
		t.Errorf("test failed")
	}
	if expected[1] != output[1] {
		log.Println("expected : ", expected[1])
		log.Println("got : ", output[1])
		t.Errorf("test failed")
	}
	if expected[2] != output[2] {
		log.Println("expected : ", expected[2])
		log.Println("got : ", output[2])
		t.Errorf("test failed")
	}

}

func TestAuthenticate(t *testing.T) {
	db, err := sql.Open("postgres",
		testConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		user := []interface{}{"David Noel", "pass3", "davidnoel@geologix.com", "PM"}
		initDB()
		ok, _ := Authenticate("some_random_user_name", "passwd", db)
		if ok {
			t.Errorf("test failed")
		}
		if InsertIntoDb(user, User{}, "users", db) == nil {
			ok, _ := Authenticate("davidnoel@geologix.com", "pass3", db)
			if !ok {
				t.Errorf("test failed")
			}
		}

	}
}

func TestInsertIntoDb(t *testing.T) {
	data := []interface{}{"PASTURIA", "description", "BUG", 3, 2, "OPEN", 3}
	var err error
	db, err := sql.Open("postgres",
		testConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		initDB()
		if InsertIntoDb(data, Issue{}, "issues", db) != nil {
			t.Errorf("test failed")
		}
		if InsertIntoDb(data, Issue{}, "issue", db) == nil {
			t.Errorf("test failed")
		}
	}

}

func TestGetRole(t *testing.T) {
	username := "some_random_user_name"
	var err error
	db, err := sql.Open("postgres",
		testConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		initDB()
		_, _, err := GetRole(db, username)
		if err == nil {
			t.Errorf("test failed")
		}
	}

}

func TestGetSampleJSONData(t *testing.T) {

	if string(GetSampleJSONData("x")) == "x" {
		t.Errorf("test failed")
	}

}

func TestInsertFromJSON(t *testing.T) {
	db, err := sql.Open("postgres",
		testConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		initDB()
		err = insertUserFromJSON("./json/users404.json", db)
		if err == nil {
			t.Errorf("test failed")
		}
		err = insertUserFromJSON("./json/users.json", db)
		if err != nil {
			t.Errorf("test failed")
		}
		err = insertProjectFromJSON("./json/projects.json", db)
		if err != nil {
			t.Errorf("test failed")
		}
		err = insertIssueFromJSON("./json/issues.json", db)
		if err != nil {
			t.Errorf("test failed")
		}

	}

}

func TestGetOpenIssues(t *testing.T) {
	db, err := sql.Open("postgres",
		testConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		issue := []interface{}{"PASTURIA", "description", "BUG", 3, 2, "OPEN", 3}
		initDB()
		issues, errr := GetOpenIssues(db)

		if errr != nil || len(issues) > 0 {
			log.Println(errr)
			t.Errorf("test failed")
		}
		if InsertIntoDb(issue, Issue{}, "issues", db) == nil {
			issues, err = GetOpenIssues(db)
			if err != nil || len(issues) != 1 {
				log.Println(err)
				t.Errorf("test failed")
			}
		}

	}
}

func TestGetEmail(t *testing.T) {
	db, err := sql.Open("postgres",
		testConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		user := []interface{}{1, "David Noel", "pass3", "davidnoel@geologix.com", "PM"}
		initDB()
		email, errr := GetEmail(db, 1)

		if errr == nil {
			log.Println(errr)
			t.Errorf("test failed")
		}
		if InsertIntoDb(user, TestUser{}, "users", db) == nil {
			email, err = GetEmail(db, 1)
			if err != nil || email != "davidnoel@geologix.com" {
				log.Println(err)
				t.Errorf("test failed")
			}
		}

	}

}

func TestGetName(t *testing.T) {
	db, err := sql.Open("postgres",
		testConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		user := []interface{}{1, "David Noel", "pass3", "davidnoel@geologix.com", "PM"}
		initDB()
		name, errr := GetName(db, 1)

		if errr == nil {
			log.Println(errr)
			t.Errorf("test failed")
		}
		if InsertIntoDb(user, TestUser{}, "users", db) == nil {
			name, err = GetName(db, 1)
			if err != nil || name != "David Noel" {
				log.Println(err)
				t.Errorf("test failed")
			}
		}

	}

}

func TestGetWatchers(t *testing.T) {

	db, err := sql.Open("postgres",
		testConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		issue := []interface{}{1, "PASTURIA", "description", "BUG", 3, 2, "OPEN", 3}
		initDB()
		watchers, errr := GetWatchers(db, 1, 3, 2)

		if errr != nil || len(watchers) != 2 {
			log.Println(errr)
			t.Errorf("test failed")
		}
		if InsertIntoDb(issue, TestIssue{}, "issues", db) == nil {
			watchers, err = GetWatchers(db, 1, 3, 2)
			if err != nil || len(watchers) != 2 {
				log.Println(err)
				t.Errorf("test failed")
			}
		}

	}
}
