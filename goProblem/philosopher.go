package main

import (
	"fmt"
	"time"
)

type Fork struct {
	id    uint
	clean bool
}

type Philosopher struct {
	id         uint
	leftFork   *Fork
	rightFork  *Fork
	leftIn     chan *Fork
	leftOut    chan *Fork
	rightIn    chan *Fork
	rightOut   chan *Fork
	shutDown   chan bool
	eatenMeals chan uint
}

func RunPhilosopher(philo Philosopher) {
	eatenMeals := uint(0)

	for {

		fmt.Printf("Philosopher %v is communicating\n", philo.id)
		if !communicateWithNeighbours(&philo) {
			break
		}

		if philo.leftFork != nil && philo.rightFork != nil {
			fmt.Printf("Philosopher %v is eating with forks: %v %v\n",
				philo.id, philo.leftFork.id, philo.rightFork.id)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Philosopher %v finished eating with forks: %v %v\n",
				philo.id, philo.leftFork.id, philo.rightFork.id)
			philo.leftFork.clean = false
			philo.rightFork.clean = false
			eatenMeals++
			select {
			case philo.eatenMeals <- eatenMeals:
			case <-philo.shutDown:
				break
			}

		}

		fmt.Printf("Philosopher %v is thinking\n", philo.id)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf("Philosopher %v is leaving\n", philo.id)
}

func communicateWithNeighbours(philo *Philosopher) bool {
	if philo.leftFork == nil && philo.rightFork == nil {
		fmt.Printf("Philosopher %v has no forks\n", philo.id)
		select {
		case philo.leftFork = <-philo.leftIn:
			philo.leftFork.clean = true
		case philo.rightFork = <-philo.rightIn:
			philo.rightFork.clean = true
		case <-philo.shutDown:
			return false
		}
	} else if philo.leftFork == nil {
		fmt.Printf("Philosopher %v has no left fork\n", philo.id)
		if philo.rightFork.clean {
			select {
			case philo.leftFork = <-philo.leftIn:
				philo.leftFork.clean = true
			case <-philo.shutDown:
				return false
			}
			return true
		}
		//else
		select {
		case philo.leftFork = <-philo.leftIn:
			philo.leftFork.clean = true
		case philo.rightOut <- philo.rightFork:
			philo.rightFork = nil
		case <-philo.shutDown:
			return false
		}

	} else if philo.rightFork == nil {
		fmt.Printf("Philosopher %v has no right fork\n", philo.id)
		if philo.leftFork.clean {
			select {
			case philo.rightFork = <-philo.rightIn:
				philo.rightFork.clean = true
			case <-philo.shutDown:
				return false
			}
			return true
		}
		//else
		select {
		case philo.rightFork = <-philo.rightIn:
			philo.leftFork.clean = true
		case philo.leftOut <- philo.leftFork:
			philo.leftFork = nil
		case <-philo.shutDown:
			return false
		}
	} else {
		if !philo.leftFork.clean {
			select {
			case philo.leftOut <- philo.leftFork:
				philo.leftOut = nil
			case <-philo.shutDown:
				return false
			default:
			}
		}
		if !philo.rightFork.clean {
			select {
			case philo.rightOut <- philo.rightFork:
				philo.rightOut = nil
			case <-philo.shutDown:
			default:
			}
		}
		select {
		case <-philo.shutDown:
			return false
		default:
		}
	}
	return true
}
