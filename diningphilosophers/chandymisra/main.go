package main

import (
	"fmt"
	"github.com/shiningrush/algorithms/diningphilosophers/chandymisra/philosopher"
)

var resource []*philosopher.Resource
var philosophers []*philosopher.Philosopher

func main(){
	initResource()
	initPhilosophers()
	for _, v := range philosophers{
		go v.Start()
	}

	g := ""
	fmt.Scanln(&g)
}

func initResource(){
	resource = append(resource ,&philosopher.Resource{})
	resource = append(resource ,&philosopher.Resource{})
	resource = append(resource ,&philosopher.Resource{})
	resource = append(resource ,&philosopher.Resource{})
	resource = append(resource ,&philosopher.Resource{})
}

func initPhilosophers(){
	resChan0 := make( chan chan *philosopher.Resource)
	resChan1 := make( chan chan *philosopher.Resource)
	resChan2 := make( chan chan *philosopher.Resource)
	resChan3 := make( chan chan *philosopher.Resource)
	resChan4 := make( chan chan *philosopher.Resource)
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph1", 1, resource[0], resource[4], resChan0, resChan4 ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph2", 2, resource[1], nil, resChan1, resChan0 ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph3", 3, resource[2], nil, resChan2, resChan1 ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph4", 4, resource[3], nil, resChan3, resChan2 ))
	philosophers = append(philosophers, philosopher.NewPhilosopher("ph5", 5, nil, nil, resChan4, resChan3 ))
}