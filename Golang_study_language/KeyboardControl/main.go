package main

import (
	"encoding/base64"
	"fmt"

	"github.com/mattn/go-tty"
)

func main() {
	// tty, err := tty.Open()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer tty.Close()

	// for {
	// 	r, err := tty.ReadRune()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// fmt.Println("Key press => " + string(r))
	// 	switch string(r) {
	// 	case "A":
	// 		fmt.Println("A")
	// 	case "B":
	// 		fmt.Println("B")
	// 	default:
	// 		// fmt.Println("Outra")
	// 	}
	// }

	tty, err := tty.Open()
	if err != nil {
		println("ERRO")
	}

	defer tty.Close()

	fmt.Print("Username: ")
	username, err := tty.ReadString()
	if err != nil {
		println("canceled")
		return
	}

	fmt.Print("Password: ")
	password, err := tty.ReadPassword()
	if err != nil {
		println("canceled")
		return
	}
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(username + ":" + password)))
}
