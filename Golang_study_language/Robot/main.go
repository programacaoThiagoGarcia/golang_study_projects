package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {

	url := "http://192.168.1.58:8080/user"
	method := "POST"

	for i := 0; i <= 100; i++ {
		age := rand.Intn(12) + 25

		name := randomString(5)

		user := User{ID: i,
			Email: name + "@gamil.com",
			Nome:  name,
			Idade: age,
		}
		b, err := json.Marshal(user)
		if err != nil {
			return
		}

		robotConnect(string(b), method, url)
		fmt.Println(name + "@gmail.com")

	}

}

func robotConnect(user string, method string, url string) {
	payload := strings.NewReader(user)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
