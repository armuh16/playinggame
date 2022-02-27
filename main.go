package main

import (
	"log"
	"time"
)

func main() {

	table := make(chan *ball)

	go player("bung", table)
	go player("arif", table)

	table <- new(ball)
	time.Sleep(1 * time.Second)
	<-table
}

type ball struct {
	hits int
}

func player(name string, table chan *ball) {
	for {
		ball := <-table
		ball.hits++
		log.Println(name, "hits the ball", ball.hits)
		time.Sleep(50 * time.Millisecond)
		table <- ball
	}
}
