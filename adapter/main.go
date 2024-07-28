package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ToDo struct {
	UserID    int    `json:"user_id"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// no adapter
	todo := GetRemoteData()
	fmt.Println("TODO without adapter:\t", todo.ID, todo.Title, todo.Completed)
}

func GetRemoteData() *ToDo {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var todo ToDo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatal(err)
	}

	return &todo
}
