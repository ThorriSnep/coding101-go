package main

import "fmt"

func main() {
	var max int
	fmt.Print("Please enter MAX: ")
	fmt.Scanln(&max)
	for i := 1; i <= max; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
	for i := max-1; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
