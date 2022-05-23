package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

//time go run main.go google.com itsector.pt github.com facebook.com twitter.com youtube.com

func main() {
	args := os.Args[1:]
	var wg sync.WaitGroup

	if len(args) < 1 {
		log.Fatalln("Use: go run main.go <url 1> <url 2> <url 3> ... <url n>")
	}

	makeRequest(args, &wg)
	// makeRequestSync(args)
	// makeRequestChan(args)

	wg.Wait()
}

func makeRequest(urls []string, wg *sync.WaitGroup) {
	var mutex sync.Mutex
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := request(url)
			if err != nil {
				log.Fatal(err)
			}
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Printf("[%v] %v \n", resp.StatusCode, url)

		}(url)
	}
}

func makeRequestSync(urls []string) {
	for _, url := range urls {
		resp, err := request(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[%v] %v \n", resp.StatusCode, url)
	}
}

func makeRequestChan(url []string) {
	ch := make(chan int)
	for _, url := range url {
		go func(url string, ch chan<- int) {
			resp, err := request(url)
			if err != nil {
				log.Fatal(err)
			}
			ch <- resp.StatusCode
		}(url, ch)
		fmt.Printf("[%v] %v \n", <-ch, url)

	}
}

func request(url string) (resp *http.Response, err error) {
	return http.Get("http://" + url)
}
