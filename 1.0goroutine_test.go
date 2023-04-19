package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() //goroutine
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
