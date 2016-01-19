// goImpl
package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 0

func thread_1() {
	for counter := 0; counter < 1000000; counter++ {
		i++

	}
}

func thread_2() {
	for counter := 0; counter < 1000000; counter++ {
		i--
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go thread_1()
	go thread_2()
	fmt.Println(i)
	fmt.Println("zzzz..")
	time.Sleep(1000 * time.Millisecond)

	fmt.Println(i)

}
