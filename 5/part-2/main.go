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
	if len(os.Args) < 2 {
		log.Fatal("too few arguments")
	}
	fmt.Println(parseInp(os.Args[1]))
}

type freshRange struct {
	start int
	end   int
}

func parseInp(filename string) int {

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	ranges := []freshRange{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		ranges = append(ranges, line2range(line))
	}

	fmt.Println(ranges)

	return 0
}

func line2range(s string) freshRange {
	fr := freshRange{}

	split := strings.Split(s, "-")

	num, err := strconv.Atoi(split[0])

	if err != nil {
		log.Fatal(err)
	}

	fr.start = num

	num, err = strconv.Atoi(split[1])

	if err != nil {
		log.Fatal(err)
	}

	fr.end = num

	if fr.start > fr.end {
		log.Fatal("(line2range): start should not be bigger than end")
	}

	return fr
}
