package routes

import (
	"fmt"
	"golangserver/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

func AddApproutes(route *mux.Router) {

	setStaticFolder(route)
	route.HandleFunc("/users/{name}", controllers.GetUsers).Methods("GET")

	fmt.Println("Routes are Loded.")
}
