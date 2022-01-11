package main

import (
	"fmt"
	"time"
)

//https://gobyexample.com/tickers
func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			fmt.Println("<=")
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	fmt.Println("=>")
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

}
