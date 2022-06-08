package main

import (
	"fmt"
	"net/http"
)

var count int = 0

func main() {
	fmt.Println("start")
	http.HandleFunc("/countup", handler)
	http.HandleFunc("/get", countHandler)
	http.ListenAndServe(":8080", nil)
	// if err := fmt.Println("shutdown"); err != nil{

	// }
}

func countHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<html><h1>Hello</h1></html>")
}

func handler(w http.ResponseWriter, r *http.Request) {

	count++
	fmt.Fprintf(w, "%v", count)
}
