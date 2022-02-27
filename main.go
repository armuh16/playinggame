package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	table := make(chan *ball)
	done := make(chan *ball)

	go player("bung", table, done)
	go player("arif", table, done)

	referree(table, done)
}

type ball struct {
	hits       int
	lastPlayer string
}

func referree(table chan *ball, done chan *ball) {
	table <- new(ball)

	for {
		select {
		case ball := <-done:
			log.Println("winner is", ball.lastPlayer)
			return
		}
	}
}

func player(name string, table chan *ball, done chan *ball) {
	for {
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		select {
		case ball := <-table:
			v := r.Intn(1000)
			if v%11 == 0 {
				log.Println(name, "drop the ball")
				done <- ball
				return
			}
			ball.hits++
			ball.lastPlayer = name
			log.Println(name, "hits the ball", ball.hits)
			time.Sleep(50 * time.Millisecond)
			table <- ball
		case <-time.After(2 * time.Second):
			return
		}
	}
}
