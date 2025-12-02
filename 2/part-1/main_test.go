package main

import (
	"fmt"
	"testing"
)

type valid_test struct {
	input           string
	expected_result bool
}

func Test_isValidID(t *testing.T) {
	tests := []valid_test{
		{"11", false},
		{"22", false},
		{"99", false},
		{"998", true},
		{"1012", true},
		{"1188511880", true},
		{"1188511890", true},
		{"222220", true},
		{"222224", true},
		{"1698522", true},
		{"1698528", true},
		{"446443", true},
		{"446449", true},
		{"38593856", true},
		{"38593862", true},
		{"565653", true},
		{"565659", true},
		{"824824821", true},
		{"824824827", true},
		{"2121212118", true},
		{"2121212124", true},
		{"1010", false},
		{"1188511885", false},
		{"222222", false},
		{"446446", false},
		{"38593859", false},
	}

	for _, test := range tests {
		result := isValidID(test.input)
		if result != test.expected_result {
			fmt.Printf("%s is %t, not %t\n", test.input, test.expected_result, result)
			t.Fail()
		}
	}
}

type getRanges_test struct {
	input             string
	expected_result_a string
	expected_result_b string
}

func Test_getRange(t *testing.T) {
	tests := []getRanges_test{
		{"11-22", "11", "22"},
		{"95-115", "95", "115"},
		{"998-1012", "998", "1012"},
		{"1188511880-1188511890", "1188511880", "1188511890"},
		{"222220-222224", "222220", "222224"},
		{"1698522-1698528", "1698522", "1698528"},
		{"446443-446449", "446443", "446449"},
		{"38593856-38593862", "38593856", "38593862"},
		{"565653-565659", "565653", "565659"},
		{"824824821-824824827", "824824821", "824824827"},
		{"2121212118-2121212124", "2121212118", "2121212124"},
	}

	for _, test := range tests {
		res_a, res_b := getRange(test.input)
		if res_a != test.expected_result_a || res_b != test.expected_result_b {
			fmt.Printf("got \"%s\" and \"%s\", instead of \"%s\" and \"%s\"\n", res_a, res_b, test.expected_result_a, test.expected_result_b)
			t.Fail()
		}
	}
}

type checkRange_test struct {
	input_a         string
	input_b         string
	expected_result int
}

func Test_checkRange(t *testing.T) {
	tests := []checkRange_test{
		{"11", "22", 11 + 22},
		{"95", "115", 99},
		{"998", "1012", 1010},
		{"1188511880", "1188511890", 1188511885},
		{"222220", "222224", 222222},
		{"1698522", "1698528", 0},
		{"446443", "446449", 446446},
		{"38593856", "38593862", 38593859},
		{"565653", "565659", 0},
		{"824824821", "824824827", 0},
		{"2121212118", "2121212124", 0},
	}

	for _, test := range tests {
		res := checkRange(test.input_a, test.input_b)
		if res != test.expected_result {
			fmt.Printf("%s, %s should give output %d, not %d\n", test.input_a, test.input_b, test.expected_result, res)
			t.Fail()
		}
	}
}

func Test_parseInput(t *testing.T) {
	res := parseInput("../test-inp")
	if res != 1227775554 {
		fmt.Printf("got %d, expected\n", 1227775554)
		t.Fail()
	}
}
