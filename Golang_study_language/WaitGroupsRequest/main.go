package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://jsonplaceholder.typicode.com/users/1",
	"https://jsonplaceholder.typicode.com/users/2",
	"https://jsonplaceholder.typicode.com/users/3",
	"https://jsonplaceholder.typicode.com/users/4",
	"https://jsonplaceholder.typicode.com/users/5",
	"https://jsonplaceholder.typicode.com/users/6",
	"https://jsonplaceholder.typicode.com/users/7",
	"https://jsonplaceholder.typicode.com/users/8",
	"https://jsonplaceholder.typicode.com/users/9",
	"https://jsonplaceholder.typicode.com/users/10",
}

func main() {
	http.HandleFunc("/users", getUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Wg controla as Goroutines
	var wg sync.WaitGroup
	// Channel traz os valores do response
	ch := make(chan http.Response, len(urls))

	connect(&wg, ch)

	resp := <-ch

	// Transforma a resposta em string
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	fmt.Fprintf(w, "%v", newStr)

	wg.Wait()
}

func connect(wg *sync.WaitGroup, ch chan<- http.Response) {
	for _, url := range urls {
		wg.Add(1)
		fmt.Println(url)
		go requestURL(url, ch, wg)
	}

}

func requestURL(url string, ch chan<- http.Response, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("==>")
	resp, err := http.Get(url)
	fmt.Println("<==")
	if err != nil {
		log.Fatal(err)
	}
	ch <- *resp
}
