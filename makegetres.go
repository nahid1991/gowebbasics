package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Todos struct {
	UserId int
	ID     int
	Title  string
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8000", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("https://jsonplaceholder.typicode.com/todos/")
	bytes, _ := ioutil.ReadAll(resp.Body)
	todos := make([]Todos, 0)
	resp.Body.Close()

	json.Unmarshal(bytes, &todos)
	res, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
