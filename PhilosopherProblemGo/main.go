package main

import (
	"fmt"
	"time"
)

func main() {
	numOfPhils := 5
	var numOfMeals uint = 5
	forks := make([]Fork, numOfPhils)
	philos := make([]Philosopher, numOfPhils)
	waiter := Waiter{make([]chan int, numOfPhils), make([]chan bool, numOfPhils), make([]*Fork, numOfPhils), make(chan bool)}
	//eatenMeals := make([]uint, numOfPhils)

	for i := 0; i < numOfPhils; i++ {
		forks[i] = Fork{i, false}
		waiter.forks[i] = &forks[i]
	}

	for i := 0; i < numOfPhils; i++ {
		philos[i] = Philosopher{i, &forks[i], &forks[(i+1)%numOfPhils], make(chan bool), make(chan int), make(chan bool), make(chan uint)}
		waiter.inputChannels[i] = philos[i].waiterIn
		waiter.outputChannels[i] = philos[i].waiterOut
	}

	go RunWaiter(waiter)

	for i := 0; i < numOfPhils; i++ {
		go RunPhilosopher(philos[i])
	}

	for {
		time.Sleep(100)
		finished := true

		for i := 0; i < numOfPhils; i++ {
			temp := <-philos[i].eatenMeals
			if temp < numOfMeals {
				finished = false
			}
			/*
				select {
				case temp := <-philos[i].eatenMeals:
					fmt.Println("chuj")
					eatenMeals[i] = temp
					if eatenMeals[i] < numOfMeals {
						finished = false
						fmt.Println(eatenMeals[i])
					}
				default:
					fmt.Println("dupa")
				}*/
		}
		if finished {
			break
		}
	}

	for i := 0; i < numOfPhils; i++ {
		philos[i].shutDown <- true
	}

	waiter.shutDown <- true

	fmt.Println("Exiting")
}
