package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(parseInp(os.Args[1]))
	} else {
		log.Fatal("you need to supply an argument")
	}
}

func parseInp(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	//find the initial placement of the beam
	scanner.Scan()
	firstLine := scanner.Text()
	index := strings.Index(firstLine, "S")
	fmt.Println("beam @ index", index)
	beams := initBeams(index, len(firstLine))
	//printBeams(beams)

	if len(beams) != len(firstLine) {
		log.Fatal("should be the same length", len(beams), len(firstLine))
	}

	for scanner.Scan() {
		line := scanner.Text()
		for i, char := range line {
			if char == '^' {
				if beams[i] != 0 {
					beams[i+1] += beams[i]
					beams[i-1] += beams[i]
					beams[i] = 0
				}
			}
		}
		//printBeams(beams)
	}

	return getPathSum(beams)
}

func getPathSum(beams []int) int {
	accum := 0
	for _, i := range beams {
		accum += i
	}
	return accum
}

func initBeams(initIndex, length int) []int {
	beams := []int{}

	for i := 0; i < length; i++ {
		init := 0
		if i == initIndex {
			init = 1
		}
		beams = append(beams, init)
	}

	return beams
}

func printBeams(beams []int) {
	for _, beam := range beams {
		if beam != 0 {
			fmt.Printf("%d", beam)
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()
}
