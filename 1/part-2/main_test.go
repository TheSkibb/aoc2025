package main

import (
	"fmt"
	"testing"
)

func Test_parse_input(t *testing.T) {
	count := parseInput("../test-inp")
	if count != 6 {
		fmt.Printf("should be %d, not %d\n", 6, count)
		t.Fail()
	}
}

type test struct {
	a               int
	b               int
	expected_result int
}

func Test_plus(t *testing.T) {
	tests := []test{
		{1, 2, 3},
		{99, 1, 0},
		{99, 2, 1},
	}

	for _, test := range tests {
		result := plus(test.a, test.b)
		if result != test.expected_result {
			fmt.Printf("%d+%d != %d (should be %d)", test.a, test.b, result, test.expected_result)
		}
	}
}

func Test_minus(t *testing.T) {
	tests := []test{
		{1, 1, 0},
		{0, 1, 99},
		{0, 100, 0},
	}

	for _, test := range tests {
		result := minus(test.a, test.b)
		if result != test.expected_result {
			fmt.Printf("%d-%d != %d (should be %d)", test.a, test.b, result, test.expected_result)
			t.Fail()
		}
	}
}
