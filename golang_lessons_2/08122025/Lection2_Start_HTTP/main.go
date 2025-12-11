package main

import (
	"fmt"
	"net/http"
	"log"
)

func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, i'm new web-server! <h1>")
	fmt.Println("r.Method: ", r.Method)
	fmt.Println("r.URL: ", r.URL)
}

func main() {
	http.HandleFunc("GET /", GetGreet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}