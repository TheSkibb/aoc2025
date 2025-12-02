package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(parseInput("input"))
}

func plus(a, b int) int {
	return (a + b) % 100
}

func minus(a, b int) int {
	return (((a - b) % 100) + 100) % 100
}

func parseInput(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("could not open file", err)
	}
	scanner := bufio.NewScanner(file)

	count := 0
	pointer := 50
	for scanner.Scan() {
		currline := scanner.Text()

		inpNum, err := strconv.Atoi(currline[1:])

		if err != nil {
			log.Fatal("could not convert string to number")
		}

		fmt.Print(pointer)

		if currline[0] == 'R' {
			pointer = minus(pointer, inpNum)
			fmt.Print(" - ", inpNum)
		}
		if currline[0] == 'L' {
			pointer = plus(pointer, inpNum)
			fmt.Print(" + ", inpNum)
		}
		fmt.Println("=", pointer)
		if pointer == 0 {
			count++
		}
	}
	return count
}
