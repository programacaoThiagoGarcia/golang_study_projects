package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", getIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getIndex(resp http.ResponseWriter, req *http.Request) {
	log.Printf("%v", req.RequestURI)

}
