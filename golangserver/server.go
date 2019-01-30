package main

import (
	"fmt"
	"net/http"

	"golangserver/routes"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server will start at port: 8000")
	//db.ConnectDatabase()
	route := mux.NewRouter()

	routes.AddApproutes(route)

	fmt.Println(http.ListenAndServe(":8000", route))
}
