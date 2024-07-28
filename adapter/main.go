package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ToDo struct {
	UserID    int    `json:"user_id" xml:"user_id"`
	ID        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

type DataInterface interface {
	GetData() (*ToDo, error)
}

// wrapper type
type RemoteService struct {
	Remote DataInterface
}

func (r *RemoteService) CallRemoteService() (*ToDo, error) {
	return r.Remote.GetData()
}

type JSONBackend struct{}

func (jb *JSONBackend) GetData() (*ToDo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var todo ToDo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

type XMLBackend struct{}

func (xb *XMLBackend) GetData() (*ToDo, error) {
	xmlFile := `
<?xml version="1.0" encoding="UTF-8"?>
<root>
	<user_id>1</user_id>
	<id>1</id>
	<title>delectus aut autem</title>
	<completed>false</completed>
</root>
`
	var toDo ToDo
	err := xml.Unmarshal([]byte(xmlFile), &toDo)
	if err != nil {
		return nil, err
	}
	return &toDo, nil
}

func main() {
	// no adapter
	todo := GetRemoteData()
	fmt.Println("TODO without adapter:\t", todo.ID, todo.Title, todo.Completed)

	// with adapter, using JSON
	jsonBackend := &JSONBackend{}
	jsonAdapter := &RemoteService{Remote: jsonBackend}
	tbFromJSON, _ := jsonAdapter.CallRemoteService()

	// with adapter, using XML
	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}
	tbFromXML, err := xmlAdapter.CallRemoteService()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("TODO from JSON adapter:\t", tbFromJSON.ID, tbFromJSON.Title, tbFromJSON.Completed)
	fmt.Println("TODO from XML adapter:\t", tbFromXML.ID, tbFromXML.Title, tbFromXML.Completed)
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
