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
	}()

	// ETL
	go func () {
		for {
			x, ok := <- naturals
			if ok == false {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		fmt.Println(<- squares)
	}
}
