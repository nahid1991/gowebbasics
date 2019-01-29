package main

import (
	"encoding/json"
	"net/http"
)

type Test struct {
	UserId int
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8001", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		var t Test
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		res, err := json.Marshal(t)

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}
