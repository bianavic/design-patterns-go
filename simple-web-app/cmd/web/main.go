package main

import (
	"fmt"
	"net/http"
)

const port = ":4000"

type application struct {
}

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello, Stranger!")
	})

	fmt.Println("Server running on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		//log.Panic(err)
		fmt.Println("Error starting server:", err)
		return
	}

}
