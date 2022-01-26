package main

import "fmt"

func main() {

	for num := 0; num < 10; num++ {
		if num%2 == 0 {
			fmt.Printf("\n%d is even", num)
		} else {
			fmt.Printf("\n%d is odd", num)
		}

	}
}
