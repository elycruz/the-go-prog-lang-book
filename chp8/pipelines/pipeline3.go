package main

import (
	"flag"
	"fmt"
)

func counter(limit int, out chan<- int) {
	for x := 0; x <= limit; x++ {
		out <- x
	}
	close(out)
}

func squarer(incoming <-chan int, out chan<- int) {
	for x := range incoming {
		out <- x * x
	}
	close(out)
}

func printer(incoming <-chan int) {
	for x := range incoming {
		fmt.Println(x)
	}
}

func main() {
	limit := flag.Int("l", -1, "Squares limit")

	flag.Parse()

	if *limit <= 0 {
		fmt.Printf("Only positive limits allowed.  %d was given.\n", limit)
		return
	}

	naturals := make(chan int)
	squares := make(chan int)

	go counter(100, naturals)

	go squarer(naturals, squares)

	printer(squares)
}
