package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Test struct {
	UserId int
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", foo)
	http.ListenAndServe(":8001", mux)
}

func foo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		var t Test
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t.UserId)
	}
}
