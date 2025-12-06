package main

import (
	"fmt"
	"testing"
)

type getOverlap_test struct {
	r1  freshRange
	r2  freshRange
	res int
}

func Test_getOverlap(t *testing.T) {
	tests := []getOverlap_test{
		{freshRange{1, 5}, freshRange{4, 7}, 2},
		{freshRange{1, 5}, freshRange{6, 9}, 0},
		{freshRange{1, 5}, freshRange{7, 100}, 0},
		{freshRange{0, 5}, freshRange{2, 3}, 2},
		{freshRange{5, 6}, freshRange{1, 4}, 0},
		{freshRange{4, 6}, freshRange{1, 4}, 1},
	}

	for _, test := range tests {
		result := getOverLap(test.r1, test.r2)
		if result != test.res {
			fmt.Printf("%d->%d, %d->%d expected overlap %d, got %d\n", test.r1[0], test.r1[1], test.r2[0], test.r2[1], test.res, result)
			t.Fail()
		}
	}
}
