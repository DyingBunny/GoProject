package main

import (
	"fmt"
	"time"
)
var done chan bool
func HelloWorld() {
	fmt.Println("Hello world goroutine")
	time.Sleep(1*time.Second)
	//done <- true
}
func main() {
	done = make(chan bool)  // 创建一个channel
	go HelloWorld()
	//<-done
}