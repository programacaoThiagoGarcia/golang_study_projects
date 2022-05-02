package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Thiago")
	fmt.Println(message)
}
