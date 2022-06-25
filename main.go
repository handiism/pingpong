package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	ball := make(chan *Ball)
	droppedBall := make(chan *Ball)

	go Player("handiism", ball, droppedBall)
	go Player("netashad", ball, droppedBall)

	Referree(ball, droppedBall)
}

type Ball struct {
	LastPlayer string
	Hits       int
}

func Referree(ball chan *Ball, droppedBall chan *Ball) {
	ball <- new(Ball)

	for {
		select {
		case currentBall := <-droppedBall:
			log.Printf("%s win the match", currentBall.LastPlayer)
			return
		}
	}
}

func Player(name string, ball chan *Ball, dropppedBall chan *Ball) {
	for {
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		select {
		case currentBall := <-ball:
			v := r.Intn(1000)

			if v%6 == 0 {
				log.Printf("%s drop the ball", name)
				dropppedBall <- currentBall
				return
			}

			currentBall.LastPlayer = name
			currentBall.Hits++
			log.Printf("%s hits the ball: %d", name, currentBall.Hits)
			time.Sleep(50 * time.Millisecond)
			ball <- currentBall
		case <-time.After(2 * time.Second):
			return
		}
	}
}
