package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "gandhi"
		channel <- "setyawan"
		channel <- "iwan"
	}()

	fmt.Println("Panjang buffer", cap(channel))            //melihat panjang buffe
	fmt.Println("Jumlah dat pada buffer : ", len(channel)) //melihat jumlah data pada buffe

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

	}()

	time.Sleep(3 * time.Second)
}
