package main

import (
	"fmt"
	"time"
)

type Fork struct {
	id   int
	used bool
}

type Philosopher struct {
	id         int
	leftFork   *Fork
	rightFork  *Fork
	waiterOut  chan bool
	waiterIn   chan int
	shutDown   chan bool
	eatenMeals chan uint
}

type Waiter struct {
	inputChannels  []chan int
	outputChannels []chan bool
	forks          []*Fork
	shutDown       chan bool
}

func RunPhilosopher(phil Philosopher) {
	var eatenTimes uint = 0
	for {
		canEat := false
		//Request forks
		phil.waiterIn <- 1
		canEat = <-phil.waiterOut
		if canEat {
			fmt.Printf("Philosoph %v started eating with forks: %v %v\n", phil.id, phil.leftFork.id, phil.rightFork.id)
			time.Sleep((500))
			fmt.Printf("Philosoph %v finished eating with forks: %v %v\n", phil.id, phil.leftFork.id, phil.rightFork.id)
		}
		//Finished eating
		phil.waiterIn <- 2
		eatenTimes += 1
		phil.eatenMeals <- eatenTimes
		select {
		case <-phil.shutDown:
			fmt.Printf("Philosopher %v finished their meal\n", phil.id)
			break
		default:
			fmt.Println("Philosopher is contemplating")
			time.Sleep((500))
		}
	}
}

func RunWaiter(waiter Waiter) {
	for {
		select {
		case <-waiter.shutDown:
			fmt.Println("Waiter shuting down")
		default:
			fmt.Println("Waiter checking for requests")
			time.Sleep(500)
		}
		for index, value := range waiter.inputChannels {
			select {
			case temp := <-value:
				if temp == 1 {
					if waiter.forks[index].used || waiter.forks[(index+1)%len(waiter.forks)].used {
						waiter.outputChannels[index] <- false
					} else {
						waiter.forks[index].used = true
						waiter.forks[(index+1)%len(waiter.forks)].used = true
						waiter.outputChannels[index] <- true
					}
				} else {
					waiter.forks[index].used = false
					waiter.forks[(index+1)%len(waiter.forks)].used = false
				}
			default:

			}
		}
	}
}
