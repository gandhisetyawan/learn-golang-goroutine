package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "gandhi setyawan"
	}()

	data := <-channel
	fmt.Println(data)
	close(channel)

}
