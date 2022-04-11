package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x++
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &m)
	}
	wg.Wait()
	fmt.Println(x)
}
