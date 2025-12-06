package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("you need to supply at least 1 argument")
	}
	fmt.Println(parseInput(os.Args[1]))
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
		line := scanner.Text()
		num, err := strconv.Atoi(line[1:])
		fmt.Printf("num %d, pointer %d\n", num, pointer)

		if err != nil {
			log.Fatal(err)
		}

		if line[0] == 'L' {
			for i := 0; i < num; i++ {
				pointer -= 1
				//fmt.Printf("pointer -> %d\n", pointer)
				if pointer == 0 {
					count++
				}
				if pointer < 0 {
					pointer += 100
				}
			}
		} else {
			for i := 0; i < num; i++ {
				pointer += 1
				if pointer >= 100 {
					pointer -= 100
				}
				if pointer == 0 {
					count++
				}
			}
		}

		fmt.Printf("pointer %d, count %d\n", pointer, count)

		fmt.Println()

	}

	return count
}
