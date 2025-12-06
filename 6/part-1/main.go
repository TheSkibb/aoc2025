package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println("result", parseInput("../test-inp"))
	fmt.Println("result", parseInput("../input"))
}

type problem struct {
	numbers  []int
	operator string
}

type problems struct {
	length int
	p      []problem
}

func parseInput(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	problems := problems{}

	//get the first line
	scanner.Scan()
	line := normalizeSpaces(scanner.Text())
	fmt.Printf("\"%s\"\n", line)
	splits := strings.Split(line, " ")
	for _, split := range splits {
		fmt.Printf("\"%s\"\n", split)
		num, err := strconv.Atoi(split)
		if err != nil {
			log.Fatal(err)
		}
		problems.p = append(problems.p, problem{
			[]int{num}, "0",
		})
	}
	problems.length = len(problems.p)

	for scanner.Scan() {
		parseLine(scanner.Text(), &problems)
	}

	accum := 0

	for _, p := range problems.p {
		res := solveProblem(p)
		fmt.Printf("problem result: %d\n", res)
		accum += res
	}

	return accum
}

func normalizeSpaces(line string) string {
	//all ocurrences of multiple spaces
	re := regexp.MustCompile(" {2,}")
	return strings.TrimSpace(re.ReplaceAllString(line, " "))
}

func parseLine(line string, problems *problems) {
	splits := strings.Split(normalizeSpaces(line), " ")

	if len(splits) != problems.length {
		log.Fatal("too few in split")
	}

	if splits[0] == "+" || splits[0] == "*" {
		for i, split := range splits {
			problems.p[i].operator = split
		}
		return
	}

	for i, split := range splits {
		num, err := strconv.Atoi(split)
		if err != nil {
			log.Fatal(err)
		}
		problems.p[i].numbers = append(problems.p[i].numbers, num)
	}
}

func solveProblem(p problem) int {

	fmt.Println("solving problems", p.numbers, "operator", p.operator)
	accum := 0
	if p.operator == "*" {
		accum = 1
		for i, num := range p.numbers {
			accum *= num
			fmt.Printf("[%d] => %d\n", i, accum)
		}
	} else if p.operator == "+" {
		for _, num := range p.numbers {
			accum += num
		}
	} else {
		log.Fatal("unknown operator")
	}
	fmt.Printf("returning %d\n", accum)

	return accum
}
