package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzbuzz(n int) {
	// Your code here!
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
