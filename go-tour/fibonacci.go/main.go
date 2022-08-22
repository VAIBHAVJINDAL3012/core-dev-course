package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	var lastButOne, last int = 0, 1

	return func() int {
		temp := lastButOne
		lastButOne = last
		last = temp + last

		return last
	}

}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
