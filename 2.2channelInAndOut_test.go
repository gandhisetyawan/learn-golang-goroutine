package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// mengirim data channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "data channel learn channel-in-and-out"

}

// menerima data channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}
