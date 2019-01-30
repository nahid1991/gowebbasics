package controller

import (
	"encoding/json"
	"fmt"
	_ "golangserver/db"
	"net/http"

	"github.com/gorilla/mux"
)

// User is Interface for user details.
type User struct {
	ID     int
	Name   string
	Email  string
	Number string
	Date   string
}

// func renderHome(response http.ResponseWriter, request *http.Request) {
// 	http.ServeFile(response, request, "views/index.html")
// }

func GetUsers(response http.ResponseWriter, request *http.Request) {
	var (
		user  User
		users []User
	)

	username := mux.Vars(request)["name"]

	rows, err := db.Query("SELECT * FROM users where name like '%" + username + "%'")

	if err != nil {
		fmt.Println(err)
		returnErrorResponse(response, request)
	}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Number, &user.Date)
		users = append(users, user)
	}
	defer rows.Close()

	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		returnErrorResponse(response, request)
	}

	if jsonResponse == nil {
		returnErrorResponse(response, request)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request) {
	jsonResponse, err := json.Marshal("It's not you it's me.")
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusInternalServerError)
	response.Write(jsonResponse)
}
