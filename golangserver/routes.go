package routes

import (
	"fmt"
	"golangserver/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

// AddApproutes will add the routes for the application
func AddApproutes(route *mux.Router) {

	setStaticFolder(route)

	// route.HandleFunc("/", renderHome)

	route.HandleFunc("/users/{name}", controller.GetUsers).Methods("GET")

	fmt.Println("Routes are Loded.")
}
