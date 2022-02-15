package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	messageArray := []string{
		"No one is perfect - that’s why pencils have erasers. - Wolfgang Riebe",
		"Have no fear of perfection - you'll never reach it. - Salvador Dali",
		"The tallest mountain started as a stone. - One Punch Man Intro",
		"Make it work. Make it nice. Make it fast. Always obey this order! - kiraa",
		"A good programmer is someone who always looks both ways before crossing a one-way street. – Doug Linder",
		"If debugging is the process of removing software bugs, then programming must be the process of putting them in. - Edsger W. Dijkstra",
	}
	var randArray = make([]int, 0, 6)
	for {
		fmt.Print("Press enter for next message: ")
		fmt.Scanln()
		if len(randArray) > 0 {
			fmt.Println(messageArray[randArray[0]])
			randArray[0] = randArray[len(randArray)-1]
			randArray = randArray[:len(randArray)-1]
		} else {
			currtime := time.Now().Unix()
			rand.Seed(currtime)
			// fill randArray with indixes from messageArray
			for i, _ := range messageArray {
				randArray = append(randArray, i)
			}
			rand.Shuffle(len(randArray), func(i, j int) { randArray[i], randArray[j] = randArray[j], randArray[i] })
			fmt.Println(messageArray[randArray[0]])
			randArray[0] = randArray[len(randArray)-1]
			randArray = randArray[:len(randArray)-1]
		}
		fmt.Print("\n")
	}
}
