package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(parseInp("../input"))
}

func parseInp(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		total += getJoltage(scanner.Text())
	}
	return total
}

func getJoltage(inp string) int {
	firstValue, firstIndex := getMax([]byte(inp), true)
	secondValue, _ := getMax([]byte(inp[firstIndex+1:]), false)

	return firstValue*10 + secondValue
}

func byte2int(b byte) int {
	return int(b) - 48
}

func getMax(inp []byte, first bool) (int, int) {
	var maxIdx, maxVal int

	for i, numbyte := range inp {
		if first && i == len(inp)-1 {
			break
		}
		num := byte2int(numbyte)
		if num > maxVal {
			maxIdx = i
			maxVal = num
		}
	}

	return maxVal, maxIdx
}
