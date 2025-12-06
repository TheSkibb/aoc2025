package main

import (
	"fmt"
	"testing"
)

type solve_test struct {
	up  problem
	res int
}

/*
123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
*/

func Test_solve(t *testing.T) {
	tests := []solve_test{
		{problem{"64", "23", "314", "+"}, 1058},
		{problem{"328", "64", "98", "*"}, 3253600},
		{problem{"51", "387", "215", "+"}, 625},
		{problem{"64", "23", "314", "*"}, 8544},
	}

	for _, test := range tests {
		res := test.up.solve()
		if res != test.res {
			fmt.Println("expected", test.res, "got", res)
			t.Fail()
		}
	}

}
