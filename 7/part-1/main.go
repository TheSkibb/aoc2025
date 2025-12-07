package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println(parseInp(os.Args[1]))
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
	printBeams(beams)

	if len(beams) != len(firstLine) {
		log.Fatal("should be the same length", len(beams), len(firstLine))
	}

	accum := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, char := range line {
			if char == '^' {
				if beams[i] {
					accum++
					beams[i] = false
					//no checks for out of bounds
					beams[i+1] = true
					beams[i-1] = true
				}
			}
		}
		printBeams(beams)
	}

	return accum
}

func initBeams(initIndex, length int) []bool {
	beams := []bool{}

	for i := 0; i < length; i++ {
		beams = append(beams, i == initIndex)
	}

	return beams
}

func printBeams(beams []bool) {
	for _, beam := range beams {
		if beam {
			fmt.Printf("|")
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()
}
