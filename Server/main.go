package main

import (
	"fmt"
	"net/http"
)

var count int = 0

func main() {
	fmt.Println("start")
	http.HandleFunc("/", handler)
	http.HandleFunc("/get", countHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("shutdown")
}

func countHandler(w http.ResponseWriter, r *http.Request) {

	count++
	fmt.Fprintf(w, "%v", count)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", count)
}
