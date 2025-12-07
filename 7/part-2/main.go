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
	printBeams(beams, firstLine)

	if len(beams) != len(firstLine) {
		log.Fatal("should be the same length", len(beams), len(firstLine))
	}

	accum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := false
		for i, char := range line {
			if char == '^' {
				if beams[i] {
					split = true
					beams[i] = false
					beams[i+1] = true
					beams[i-1] = true
				}
			}
		}
		if split {
			accum += getNumberOfBeams(beams)
		}
		printBeams(beams, line)
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

func getNumberOfBeams(beams []bool) int {
	accum := 0

	for _, beam := range beams {
		if beam {
			accum++
		}
	}

	return accum
}

func printBeams(beams []bool, line string) {
	for i, beam := range beams {
		if line[i] == '^' {
			fmt.Printf("^")
		} else if beam {
			fmt.Printf("|")
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()
}
