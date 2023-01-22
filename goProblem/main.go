package main

import (
	"fmt"
	"time"
)

func main() {
	numOfPhils := uint(5)
	var numOfMeals uint = 5
	forks := make([]Fork, numOfPhils)
	philos := make([]Philosopher, numOfPhils)
	eatenMeals := make([]uint, numOfPhils)

	for i := uint(0); i < numOfPhils; i++ {
		forks[i] = Fork{i, false}
	}

	for i := uint(0); i < numOfPhils; i++ {
		philos[i] = Philosopher{i, &forks[i], nil, make(chan *Fork), make(chan *Fork), nil, nil, make(chan bool), make(chan uint)}
	}

	for i := uint(0); i < numOfPhils; i++ {
		philos[i].rightIn = philos[(i+1)%numOfPhils].leftOut
		philos[i].rightOut = philos[(i+1)%numOfPhils].leftIn
	}

	for i := uint(0); i < numOfPhils; i++ {
		go RunPhilosopher(philos[i])
	}

	for {
		time.Sleep(100 * time.Millisecond)

		for i := uint(0); i < numOfPhils; i++ {
			select {
			case temp := <-philos[i].eatenMeals:
				eatenMeals[i] = temp
			default:
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
		finished := true
		for i := uint(0); i < numOfPhils; i++ {
			if eatenMeals[i] < numOfMeals {
				finished = false
			}
		}
		if finished {
			break
		}
	}

	for i := uint(0); i < numOfPhils; i++ {
		philos[i].shutDown <- true
	}

	fmt.Println("Exiting")
}
