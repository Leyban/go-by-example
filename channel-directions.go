package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func pang(pong <-chan string, pang chan<- string) {
	msg := <-pong
	pang <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	pangs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	pang(pongs, pangs)
	fmt.Println(<-pangs)
}
