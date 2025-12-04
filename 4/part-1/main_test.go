package main

import (
	"fmt"
	"testing"
)

func Test_parseInput(t *testing.T) {
	res := parseInput("../test-inp")
	expected := 13

	if res != expected {
		fmt.Printf("got %d, expected %d\n", res, expected)
		t.Fail()
	}
}

type checkAdjecent_test struct {
	inp_pg paperGrid
	inp_x  int
	inp_y  int
	res    int
}

func Test_checkAdjecent(t *testing.T) {
	testBoard := paperGrid{
		{'.', '.', '@', '@', '.', '@', '@', '@', '@', '@'},
		{'@', '@', '@', '.', '@', '.', '@', '.', '@', '@'},
		{'@', '@', '@', '@', '@', '.', '@', '.', '@', '@'},
		{'@', '.', '@', '@', '@', '@', '.', '.', '@', '.'},
		{'@', '@', '.', '@', '@', '@', '@', '.', '@', '@'},
		{'.', '@', '@', '@', '@', '@', '@', '@', '.', '@'},
		{'.', '@', '.', '@', '.', '@', '.', '@', '@', '@'},
		{'.', '.', '.', '@', '@', '.', '@', '@', '@', '@'},
		{'.', '.', '.', '@', '@', '@', '@', '@', '@', '.'},
		{'.', '.', '.', '.', '@', '@', '@', '@', '@', '.'},
	}

	tests := []checkAdjecent_test{
		{testBoard, 0, 2, 4},
		{testBoard, 2, 2, 6},
		{testBoard, 4, 3, 7},
		{testBoard, 6, 2, 2},
		{testBoard, 7, 8, 8},
		{testBoard, 7, 9, 5},
		{testBoard, 9, 0, 3},
		//{testBoard, 1, 8, 0},
		{testBoard, 1, 6, 2},
	}

	/*
		for _, row := range testBoard {
			fmt.Println(string(row))
		}
	*/

	for _, test := range tests {
		res := checkAdjecent(test.inp_pg, test.inp_x, test.inp_y)
		if res != test.res {
			fmt.Printf("expected %d, got %d for (%d, %d)\n", test.res, res, test.inp_x, test.inp_y)
			t.Fail()
		}
	}
}

type checkSquare_test struct {
	inp_pg paperGrid
	inp_x  int
	inp_y  int
	res    bool
}

func Test_checkSquare(t *testing.T) {
	testBoard := paperGrid{
		{'.', '.', '@', '@', '.', '@', '@', '@', '@', '.'},
		{'@', '@', '@', '.', '@', '.', '@', '.', '@', '@'},
		{'@', '@', '@', '@', '@', '.', '@', '.', '@', '@'},
		{'@', '.', '@', '@', '@', '@', '.', '.', '@', '.'},
		{'@', '@', '.', '@', '@', '@', '@', '.', '@', '@'},
		{'.', '@', '@', '@', '@', '@', '@', '@', '.', '@'},
		{'.', '@', '.', '@', '.', '@', '.', '@', '@', '@'},
		{'@', '.', '@', '@', '@', '.', '@', '@', '@', '@'},
		{'.', '@', '@', '@', '@', '@', '@', '@', '@', '.'},
		{'@', '.', '@', '.', '@', '@', '@', '.', '@', '.'},
	}

	tests := []checkSquare_test{
		{testBoard, 0, 0, false},
		{testBoard, 0, 2, true},
		{testBoard, 2, 2, true},
		{testBoard, 4, 3, true},
		{testBoard, 6, 2, true},
		{testBoard, 1, 3, false},
		{testBoard, -1, 3, false},
		{testBoard, 100000, 3, false},
		{testBoard, 1, -3, false},
		{testBoard, 1, 10000, false},
	}

	/*
		for _, row := range testBoard {
			fmt.Println(string(row))
		}
	*/

	for _, test := range tests {
		res := checkSquare(test.inp_pg, test.inp_x, test.inp_y)
		if res != test.res {
			fmt.Printf("expected %t, for (%d, %d)\n", test.res, test.inp_x, test.inp_y)
			t.Fail()
		}
	}
}
