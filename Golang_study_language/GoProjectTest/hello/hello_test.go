package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestCalculate(t *testing.T) {
	want := 3
	if got := Calculate(1, 2); got != want {
		t.Error("Calculate Error")
	}
}
