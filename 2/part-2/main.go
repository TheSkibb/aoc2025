package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(parseInput(os.Args[1]))
	} else {
		log.Fatal("you need to supply the filename")
	}
}

func parseInput(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("could not open file", filename)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(splitAtComma)

	total := 0

	for scanner.Scan() {
		a, b := getRange(scanner.Text())
		total += checkRange(a, b)
	}

	return total
}

func splitAtComma(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func getRange(inp string) (string, string) {

	ranges := strings.Split(inp, "-")

	if len(ranges) != 2 {
		panic(errors.New(fmt.Sprintf("wrong length of ranges %d instead of 2", len(ranges))))
	}

	return ranges[0], ranges[1]
}

// returns the sum of the invalid ids in the range
func checkRange(start, end string) int {
	startInt, err := strconv.Atoi(strings.Trim(start, "\n"))
	if err != nil {
		log.Fatalf("could not convert \"%s\" to number: %s", start, err.Error())
	}
	endInt, err := strconv.Atoi(strings.Trim(end, "\n"))
	if err != nil {
		log.Fatalf("could not convert \"%s\" to number: %s", end, err.Error())
	}

	total := 0

	for i := startInt; i <= endInt; i++ {
		if !isValidID(strconv.Itoa(i)) {
			total += i
		}
	}
	return total
}

// checks if id is valid
// if the id is a repeating number, eg 22, 2121, 123123, 121212 it is invalid
func isValidID(inp string) bool {
	length := len(inp)

	for i := length / 2; i > 0; i-- {
		if length%i != 0 {
			continue
		}
		check := inp[:i]
		if checkRepeating(inp, check) {
			return false
		}
	}

	return true
}

// check if patt is a repeating pattern in str
func checkRepeating(str, patt string) bool {
	splits := strings.Split(str, patt)
	for _, split := range splits {
		if split != "" {
			return false
		}
	}
	return true
}
