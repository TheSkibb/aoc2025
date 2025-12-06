package main

import(
	"bufio"
	"strconv"
	"strings"
	"os"
	"log"
	"fmt"
)

type freshRange [2]int

func main(){
	fmt.Println(parseInp("../input"))
	//fmt.Println(parseInp("../test-inp"))
}

func parseInp(filename string) int {
	file, err := os.Open(filename)

	if err != nil{
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanningRanges := true

	ranges := []freshRange{}

	count := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			scanningRanges = false
			continue
		}
		if scanningRanges {
			ranges = append(ranges, getRange(line))
		}else{
			if checkId(ranges, line){
				count++
			}
		}
	}

	return count
}

func getRange(inp string) freshRange{
	split := strings.Split(inp, "-")
	if len(split) < 2 {
		log.Fatal("too few in split")
	}

	firstNum, err1 := strconv.Atoi(split[0])
	secondNum, err2 := strconv.Atoi(split[1])

	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}

	return freshRange{firstNum, secondNum}
}

func checkId(ranges []freshRange, inp string) bool{
	checkNum, err := strconv.Atoi(inp)
	if err != nil{
		log.Fatal(err)
	}
	for _, r := range ranges{
		fmt.Printf("checking %d <= %d <= %d\n", r[0], checkNum, r[1])
		if checkNum >= r[0] && checkNum <= r[1]{
			return true
		}
	}

	return false
}

