package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	table := make(chan *ball)
	done := make(chan *ball)
	go Player("ping", table, done)
	go Player("pong", table, done)

	referre(table, done)
}

type ball struct {
	hits       int
	lastPlayer string
}

func referre(table chan *ball, done chan *ball) {
	table <- new(ball)
	for {
		select {
		case ball := <-done:
			log.Println("winner is ", ball.lastPlayer)
			return
		}
	}
}

func Player(name string, table chan *ball, done chan *ball) {

	for {
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		select {
		case ball := <-table:
			v := r.Intn(1000)
			if v%11 == 0 {
				log.Println("dropping ball")
				done <- ball
				return
			}

			ball.hits++
			ball.lastPlayer = name
			log.Println(name, "ball hits:", ball.hits)
			time.Sleep(50 * time.Millisecond)
			table <- ball
		case <-time.After(2 * time.Second):
			return
		}
	}

}
