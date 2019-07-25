package main

import "fmt"

func goroutineA(a <-chan int) {
	val := <-a
	fmt.Println(val)
}

func main() {
	ch := make(chan int)
	//ch <- 3
	go goroutineA(ch)
	ch <- 3
	//ch <- 4


}

