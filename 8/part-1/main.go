package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(parseInput(os.Args[1]))
	} else {
		fmt.Println("need filename")
	}
}

type coordinate struct {
	x int
	y int
	z int
}

func parseInput(filename string) int {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
	}

	return 0
}
