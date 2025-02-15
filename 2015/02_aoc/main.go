package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"slices"
	"sort"
	// "strings"

)

func parseDimensions(ScanLine string) ([]int){
	// dimensions := strings.Split(ScanLine, "x")
	var l, w, h int
	var DimSlice []int
	_, err := fmt.Sscanf(ScanLine, "%dx%dx%d", &l, &w,&h)

	if err != nil {
		log.Fatal(err)
	}
	DimSlice = append(DimSlice, l, w, h)
	// fmt.Println(l, w, h)
	return DimSlice
}

func calculateArea(DimSlice []int) (SurfaceArea int){
	// Formula
	// 2*l*w + 2*w*h + 2*h*l
	var l, w, h int 
	l = DimSlice[0]
	w = DimSlice[1]
	h = DimSlice[2]

	array := []int{l*w, w*h, h*l}
	var MinSide int = slices.Min(array)
	
	SurfaceArea = (2*array[0]) + (2*array[1]) + (2*array[2]) + MinSide
	// fmt.Println(DimSlice, array, SurfaceArea)

	return SurfaceArea
}

func calculateRibbon(DimSlice []int) (RibbonLength int) {
	sort.Ints(DimSlice)
	s1 := DimSlice[0]
	s2 := DimSlice[1]
	s3 := DimSlice[2]
	
	RibbonWrap := s1 + s1 + s2 + s2
	RibbonBow := s1 * s2 * s3
	RibbonLength = RibbonWrap + RibbonBow
	
	return RibbonLength
}


func main() {
	// fmt.Print("Anything")
	file, err := os.Open("/Users/lf16/Documents/GitHub/Advent_of_Code/puzzle_inputs/2015/02_aoc_2015.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var sum int = 0
	var sumRibbon int = 0
	for scanner.Scan() {
		var ScanLine string = scanner.Text()

		DimSlice := parseDimensions(ScanLine)

		SurfaceArea := calculateArea(DimSlice)
		sum += SurfaceArea

		RibbonLength := calculateRibbon(DimSlice)
		sumRibbon += RibbonLength
		
	}
	fmt.Printf("Total square feet wrap: %d\nTotal square feet ribbon: %d\n", sum, sumRibbon)

}