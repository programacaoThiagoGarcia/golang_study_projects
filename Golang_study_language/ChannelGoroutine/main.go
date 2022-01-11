package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("init..")
	msgReq := request()
	fmt.Println(msgReq)
	fmt.Println("Finnish")
	fmt.Println("======================================================")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println("======================================================")

	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
	fmt.Println("======================================================")

	success := make(chan bool)
	erro := make(chan bool)
	go requestWithParameters(2, success, erro)
	select {
	case <-success:
		fmt.Println("Impar")
	case <-erro:
		fmt.Println("Par")
	}
	fmt.Println("=========== Non-Blocking Channel Operations ========")

	messages := make(chan string, 1)
	signals := make(chan string, 1)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	signals <- <-messages // Como Messages tem buffer de 1, quando envia este valor ele fica vazio, por isso não é selecionado no select

	close(signals)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	queue := make(chan string, 2)
	queue <- "0"
	queue <- "1"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}

func request() string {
	done := make(chan bool, 1)
	fmt.Println("Request")
	go worker(done)
	resp := <-done
	close(done)
	if resp {
		fmt.Println("Response")
		return "response"
	} else {
		return "NOK"
	}
}

func requestWithParameters(i int, success chan bool, erro chan bool) {
	done := make(chan bool)
	go worker(done)
	resp := <-done
	close(done)
	if resp {
		if i%2 != 0 {
			success <- true
			return
		}
		erro <- true
	}

}

func worker(done chan bool) {
	fmt.Println("Work starting...")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v \n", i)
		time.Sleep(time.Second)
	}
	fmt.Println("Work done")
	done <- true
}

// ======================================================
// 	Declaração de canais unilaterais
// r := make(<-chan type)     ## read only channel
// w := make(chan<- type)     ## write only channel
// ======================================================

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
