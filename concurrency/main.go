package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true
	close(doneChan)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	dones := make([]chan bool, 4)
	for i := range dones {
		dones[i] = make(chan bool)
	}

	go greet("Nice to meet you!", dones[0])
	go greet("How are you?", dones[1])
	go slowGreet("I .... am ..... fine....thanks", dones[2])
	go greet("why so slow?", dones[3])
	for _, done := range dones {
		<-done
	}
}
