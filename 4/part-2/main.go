package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type paperGrid [][]byte

func main() {
	if len(os.Args) > 1 {
		fmt.Println(parseInput(os.Args[1]))
	} else {
		log.Fatal("you need to supply an argument")
	}
}

func parseInput(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pg paperGrid

	for scanner.Scan() {
		line := scanner.Bytes()
		c := make([]byte, len(line))
		copy(c, line)
		pg = append(pg, c)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	removed := 0

	for {
		removedOld := removed
		removed += removeAsManyAsPossible(pg)
		if removed-removedOld == 0 {
			break
		}
	}

	return removed
}

func removeAsManyAsPossible(pg paperGrid) int {
	acc := 0
	for y := range pg {
		for x := range pg[y] {
			if pg[y][x] == '@' {
				if checkAdjecent(pg, x, y) < 4 {
					acc++
					pg[y][x] = '.'
				}
			}
		}
	}
	return acc
}

// returns true if there are less than four adjecent rolls
func checkAdjecent(pg paperGrid, x, y int) int {

	if pg[y][x] != '@' {
		log.Fatalf("checking wrong square (%d, %d)", x, y)
	}

	type check struct {
		x int
		y int
	}

	checks := []check{
		//row above
		{-1, -1},
		{-1, 0},
		{-1, +1},

		//same row
		{0, -1},
		{0, +1},

		//row below
		{+1, -1},
		{+1, 0},
		{+1, +1},
	}

	adjecent := 0
	for _, c := range checks {
		if checkSquare(pg, x+c.x, y+c.y) {
			adjecent++
		} else {
		}
	}

	return adjecent
}

func checkSquare(pg paperGrid, x, y int) bool {

	if y < 0 || x < 0 {
		return false
	}

	if y >= len(pg) {
		return false
	}
	if x >= len(pg[y]) {
		return false
	}

	if pg[y][x] == '@' {
		return true
	}

	return false
}
