package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

const passwordRegex = `^[^%_&]{8,9}$`
const emailRegex = `^[a-z]*@[a-z]*\.com$`

// Person ...
type Person struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Customer ...
type Customer struct {
	ID int
	Person
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	username, password, ok := r.BasicAuth()
	if !ok || !authenticate(username, password) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var user Person
	err := decoder.Decode(&user)
	if err != nil || !validate(user) {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	inserted, inserErr := insertCustomer(user)
	if !inserted {
		http.Error(w, inserErr.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))
}

func insertCustomer(customer Person) (bool, error) {
	_, err := db.Exec("INSERT INTO customer(username,email,password) VALUES($1,$2,$3);", customer.UserName, customer.Email, customer.Password)
	if err != nil {
		return false, err
	}
	return true, nil
}

func validate(user Person) bool {
	if len(user.UserName) < 5 || !verifyPattern(emailRegex, user.Email) || !verifyPattern(passwordRegex, user.Password) {
		return false
	}
	return true
}

func verifyPattern(pattern, actualString string) bool {
	matched, _ := regexp.MatchString(pattern, actualString)
	return matched
}
