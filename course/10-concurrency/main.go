package main

import (
	"fmt"
	"time"
)

func main(){
	// dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	// dones[1] = make(chan bool)
	// dones[2] = make(chan bool)
	// dones[3] = make(chan bool)


	// go greet("hi", dones[0])
	// go greet("how are you", dones[1])
	// go slow_greet("awesome", dones[2])
	// go greet("cool", dones[3])

	// // for _, val := range dones {
	// // 	<- val
	// // }

	// for i := 0; i<len(dones); i++ {
	// 	<- dones[i]
	// }

	done := make(chan bool)
	go greet("hi", done)
	go greet("how are you", done)
	go slow_greet("awesome", done)
	go greet("cool", done)

	for range done {
		// fmt.Println(doneChan)
	}
}

func greet(message string, doneChan chan bool){
	fmt.Println("Hello!", message)
	doneChan <- true
}

func slow_greet(message string, doneChan chan bool){
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", message)
	doneChan <- true
	close(doneChan)
}

