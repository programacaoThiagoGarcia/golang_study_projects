package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	pessoa := Pessoa{name: "Thiago", age: 41}

	hello := pessoa.sayHello()

	fmt.Println(hello)
	wg.Add(1)
	go makerequest(&wg)
	wg.Wait()

}

func makerequest(wg *sync.WaitGroup) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	defer wg.Done()

	fmt.Println("Response code : " + resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var jsonResp []Posts

	if err := json.Unmarshal(body, &jsonResp); err != nil {
		panic(err)
	}

	for index, item := range jsonResp {
		fmt.Printf("%v => %v \n", index, item.Title)
	}
}
