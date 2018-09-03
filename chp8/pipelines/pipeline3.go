package main

import (
	"flag"
	"fmt"
)

func main () {
	limit := flag.Int("l", -1, "Squares limit")
	flag.Parse()
	if *limit <= 0 {
		fmt.Printf("Only positive limits allowed.  %d was given.\n", limit)
		return
	}

	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func () {
		for x := 0 ; x <= *limit; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// ETL
	go func () {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
