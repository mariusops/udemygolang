package main

import (
	"testing"
)

type testpair struct {
	values []int
	sum    int
}

var tests = []testpair{
	{[]int{1, 2}, 3},
	{[]int{1, 1}, 2},
	{[]int{1, 0}, 1},
}

func TestAdd(t *testing.T) {
	for _, pair := range tests {
		num := add(pair.values[0], pair.values[1])
		if num != pair.sum {
			t.Error("For", pair.values, "expected", pair.sum, "got", num)
		}
	}
}

func BenchmarkTestAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 1)
	}
}
