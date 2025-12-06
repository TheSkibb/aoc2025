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
	fmt.Println("starting...")
	//fmt.Println(parseInp("../test-inp", 3))
	fmt.Println(parseInp("../input", 4))
}

type problem struct {
	index    int //index of the column in the
	operator string
	numbers  [4]string
}

func (p problem) solve(numRows int) int {

	accum := 0
	if p.operator == "*" {
		accum = 1
	}
	fmt.Println("accum", accum, "operator", p.operator)

	for i := 0; i < len(p.numbers[0]); i++ {
		numberStr := ""
		for j := 0; j < numRows; j++ {
			numberStr += string(p.numbers[j][i])
		}

		numberStr = strings.TrimSpace(numberStr)

		num, err := strconv.Atoi(numberStr)

		if err != nil {
			continue
		}

		fmt.Println(numberStr, accum)

		if p.operator == "*" {
			accum *= num
		}
		if p.operator == "+" {
			accum += num
		}
	}

	return accum
}

func parseInp(filename string, numRows int) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("getting lines")
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for i := 0; i <= numRows; i++ {
		scanner.Scan()
		if i != numRows {
			lines = append(lines, scanner.Text())
		}
	}
	//go to the fourth line
	//strings.index
	operatorLine := scanner.Text()
	problems := []problem{}
	fmt.Println(operatorLine)

	offset := 0
	for strings.IndexAny(operatorLine[offset:], "+*") != -1 {
		checkLine := operatorLine[offset:]
		index := strings.IndexAny(checkLine, "+*")
		problems = append(problems, problem{
			index:    offset + index,
			operator: string(operatorLine[offset+index]),
		})
		offset = index + offset + 1
	}

	fmt.Println("got", len(problems), "lines")

	for i, line := range lines {
		for j, _ := range problems {
			if j != len(problems)-1 {
				problems[j].numbers[i] = string(line[problems[j].index:problems[j+1].index])
			} else {
				problems[j].numbers[i] = string(line[problems[j].index:])
			}
		}
	}

	accum := 0
	for _, p := range problems {
		accum += p.solve(numRows)
	}

	return accum
}
