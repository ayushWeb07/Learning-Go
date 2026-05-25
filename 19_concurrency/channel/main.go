package main

import "fmt"

func main() {
	// create an unbuffered channel
	textCh := make(chan string)

	// fire up a goroutine & send value to channel
	go func() {
		textCh <- "Hello World XD"
	}()

	// receive value from channel
	msg := <-textCh
	fmt.Println(msg)
}
