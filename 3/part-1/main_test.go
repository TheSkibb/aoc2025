package main

import (
	"fmt"
	"testing"
)

type getJoltage_test struct {
	inp string
	res int
}

func Test_getJoltage(t *testing.T) {
	tests := []getJoltage_test{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	for _, test := range tests {
		res := getJoltage(test.inp)
		if res != test.res {
			fmt.Printf("joltage for \"%s\" is %d not %d\n", test.inp, test.res, res)
			t.Fail()
		}
	}
}

type byte2int_test struct {
	i byte
	r int
}

func Test_byte2int(t *testing.T) {
	tests := []byte2int_test{
		{'1', 1},
		{'3', 3},
		{'0', 0},
		{'9', 9},
	}

	for _, test := range tests {
		r := byte2int(test.i)
		if r != test.r {
			fmt.Printf("expected %d got %d\n", test.r, r)
			t.Fail()
		}
	}
}

type getMax_test struct {
	i   []byte
	r_1 int
	r_2 int
}

func Test_getMax(t *testing.T) {
	tests := []getMax_test{
		{[]byte("987654321111111"), 9, 0},
		{[]byte("811111111111119"), 8, 0},
		{[]byte("234234234234278"), 7, 13},
		{[]byte("818181911112111"), 9, 6},
	}

	for _, test := range tests {
		r_1, r_2 := getMax(test.i, true)
		if r_1 != test.r_1 || r_2 != test.r_2 {
			fmt.Printf("expected %d, %d got %d, %d\n", test.r_1, test.r_2, r_1, r_2)
			t.Fail()
		}
	}
}
