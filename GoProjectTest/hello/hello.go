package hello

import "rsc.io/quote"

func Hello() string {
	return quote.Hello()
}

func Calculate(num1 int, num2 int) int {
	return num1 + num2
}
