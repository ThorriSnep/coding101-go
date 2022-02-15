package main

import "fmt"

func main() {
	var loadingState int
	fmt.Print("Please enter loding state !!only int!!: ")
	fmt.Scanln(&loadingState)
	if (loadingState > 100) || (loadingState < 0) || (loadingState%10 != 0) {
		fmt.Println("Not a valide user input. Please only input variables that are in the range of [0,100] and dividable 10 without remainder")
		return
	}
	for i := loadingState / 10; i < 11; i++ {
		fmt.Print("[")
		for j := 1; j < 11; j++ {
			if j <= i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("] ", i*10, "% progress\n")
	}
}
