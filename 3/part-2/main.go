package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	offset := 0
	total := 0

	for i := 0; i < 12; i++ {
		end := len(inp) - 11 + i
		chunk := []byte(inp[offset:end])
		val, index := getMax(chunk)
		total += int(math.Pow10(11-i)) * val
		offset = index + 1 + offset
	}

	return total
}

func byte2int(b byte) int {
	return int(b) - 48
}

func getMax(inp []byte) (int, int) {
	var maxIdx, maxVal int

	for i, numbyte := range inp {
		num := byte2int(numbyte)
		if num > maxVal {
			maxIdx = i
			maxVal = num
		}
	}

	return maxVal, maxIdx
}
