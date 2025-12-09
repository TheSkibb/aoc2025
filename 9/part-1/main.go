package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(parseInp(os.Args[1]))
	} else {
		log.Fatal("supply a filname")
	}
}

type coordinate struct {
	x int
	y int
}

func parseInp(filename string) int {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	coords := []coordinate{}

	for scanner.Scan() {
		line := scanner.Text()
		coords = append(coords, getCoor(line))
	}

	fmt.Println(coords)

	return getBiggestSquare(coords)
}

func getBiggestSquare(coords []coordinate) int {

	currBiggest := 0

	for _, c1 := range coords {
		for _, c2 := range coords {
			if c1 == c2 {
				continue
			}
			fmt.Printf("{%d, %d}, {%d, %d} ", c1.x, c1.y, c2.x, c2.y)
			size := getSquareSize(c1, c2)
			fmt.Printf(" => %d", size)
			if size > currBiggest {
				fmt.Printf(" (new biggest)")
				currBiggest = size
			}
			fmt.Println()
		}
	}

	return currBiggest
}

func absolute(i int) int {
	if i < 0 {
		return -(i)
	}
	return i
}

func getSquareSize(c1, c2 coordinate) int {

	lenX := absolute(c1.x-c2.x) + 1
	lenY := absolute(c1.y-c2.y) + 1
	fmt.Printf("%d * %d ", lenX, lenY)

	return lenX * lenY
}

func getCoor(line string) coordinate {

	c := coordinate{}

	numStrs := strings.Split(line, ",")

	if len(numStrs) != 2 {
		log.Fatal("should be two")
	}

	x, err := strconv.Atoi(numStrs[0])

	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(numStrs[1])

	if err != nil {
		log.Fatal(err)
	}

	c.x = x
	c.y = y

	return c
}
