package main

import (
	"testing"
	"reflect"
)

func TestMerge(t *testing.T)  {
	var cases = []struct {
		n1 	[]int
		n2 	[]int
		m 	int
		n 	int
		exp []int
	} {
		{
			[]int{1, 3, 0, 0, 0},
			[]int{2, 4, 5},
			2,
			3,
			[]int{1, 2, 3, 4, 5},
		},
	}

	t.Run("Test order 5 numbers [1, 3] [2, 4, 5]", func(t *testing.T) {
		Merge(&cases[0].n1, cases[0].m, cases[0].n2, cases[0].n)
		if reflect.DeepEqual(cases[0].n1, cases[0].exp) == false {
			t.Errorf("Expected %v, but has %v", cases[0].exp, cases[0].n1)
		}
	})
}