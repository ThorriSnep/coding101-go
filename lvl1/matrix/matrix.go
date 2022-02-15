package main

import (
	"fmt"
	"math/rand"
	"time"
)

func print2DSlice(printSlice [][]byte) {
	/* This funktion prints an 2D byte slice in an easily readable form */
	for i, _ := range printSlice {
		for j, _ := range printSlice[0] {
			fmt.Printf("%c", printSlice[i][j])
		}
		fmt.Print("\n")
	}
}

func main() {
	/* Get all user inputs */
	var hight int
	fmt.Print("Please enter hight !!only int!!: ")
	fmt.Scanln(&hight)
	var width int
	fmt.Print("Please enter width !!only int!!: ")
	fmt.Scanln(&width)
	var nrOfRand int
	fmt.Print("Please enter amount of random numbers !!only int!!: ")
	fmt.Scanln(&nrOfRand)

	/* Verify user input */
	if nrOfRand > hight*width {
		fmt.Println("You inputed more random fields then the matrix has fields!")
		return
	}

	finalSlice := make([][]byte, hight)
	widthSlice := make([]byte, width)

	/* Create 2D Slice */
	for i := 0; i < width; i++ {
		widthSlice[i] = '#'
	}
	for i := 0; i < hight; i++ {
		tmp := make([]byte, width)
		copy(tmp, widthSlice)
		finalSlice[i] = tmp
	}

	/* Make random entries */
	var changed int
	for changed < nrOfRand {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(width)
		i := rand.Intn(hight)
		if finalSlice[i][j] != '*' {
			finalSlice[i][j] = '*'
			changed += 1
		}
	}

	print2DSlice(finalSlice)

}
