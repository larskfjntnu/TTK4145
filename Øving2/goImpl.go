// goImpl
package main

import (
	"fmt"
	"sync"
	"time"
)

var i int := 0

func thread_1(messageChannel chan int) {
	tempi := <-messageChannel
	for counter := 0; counter < 1000000; counter++ {
		tempi++
	}
	messageChannel<-tempi
}

func thread_2(messageChannel chan int) {
	tempi := <-messageChannel
	for counter := 0; counter < 1000000; counter++ {
		tempi++
	}
	messageChannel<-tempi
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	
	channel := make(chan int, 1);
	channel <- i;
	
	go thread_1(channel)
	go thread_2(channel)
	time.Sleep(1000 * time.Millisecond)
	
	i := <-channel

	fmt.Println(i)
	/* The scheduler suspends and starts the threads at random,
	meaning the first thread could increment to any number < 10^6
	then the second thread starts decrementing it to any
	> i's current value - 10^6 and so on.
	*/

}
