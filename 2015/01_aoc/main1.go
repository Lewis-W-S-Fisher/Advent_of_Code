package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	data, err := os.ReadFile("/Users/lf16/Documents/GitHub/Advent_of_Code/puzzle_inputs/2015/01_aoc_2015.txt")
	// data := aoc1read()
	if err != nil {
		log.Fatal(err)
	}

	// var length = len(data)
	var UpCount int = 0
	var DownCount int = 0
	var RunningCount int = 0 // Part 2 
	var FirstBasement bool = false

	for pos, char := range string(data) {

		if char == '(' {
			UpCount += 1
			RunningCount += 1
		}

		if char == ')' {
			DownCount += 1
			RunningCount -= 1
		}

		if RunningCount == -1 && FirstBasement == false {
			var FirstBasementPos int = pos + 1 
			fmt.Printf("First basement -1 position: %d\n", FirstBasementPos)
			FirstBasement = true
		}
		// fmt.Println(pos)

	
	}
	fmt.Printf("Up floors: %d\nDown floors: %d\n", UpCount, DownCount)

	var FloorStart int = 0
	CurrentFloor := FloorStart - DownCount + UpCount 
	fmt.Printf("Current Floor is: %d\n", CurrentFloor)
}