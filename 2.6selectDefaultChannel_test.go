package learn_golang_goroutine

import (
	"fmt"
	"testing"
)

func TestSelectDefaultChannel(t *testing.T) {
	channelData1 := make(chan string)
	channelData2 := make(chan string)
	defer close(channelData1)
	defer close(channelData2)

	go GiveMeResponse(channelData1)
	go GiveMeResponse(channelData2)

	counter := 0
	for {
		select {
		case data := <-channelData1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channelData2:
			fmt.Println("data dari channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data channel diterima")
			fmt.Println(counter)
		}
		if counter == 2 {
			break
		}
	}

}
