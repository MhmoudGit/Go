package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello world")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("server is running on port: 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
