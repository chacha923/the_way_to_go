package main

import (
	"fmt"
)

func tel_close(ch chan int) {
	for i:=0; i < 15; i++ {
		ch <- i
	}
	close(ch) // if this is ommitted: panic: all goroutines are asleep - deadlock!
}

func main() {
	var ok = true
	var i int
	ch := make(chan int)

	go tel_close(ch)
	for ok {
		if i, ok= <-ch; ok {
			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
		}
	}
}