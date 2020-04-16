package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

type input struct {
	n    int
	p, q []int
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{n: 3, p: []int{1, 3, 2}, q: []int{3, 1, 2}}, want: 3},
	{input: input{n: 8, p: []int{7, 3, 5, 4, 2, 1, 6, 8}, q: []int{3, 8, 2, 5, 4, 6, 7, 1}}, want: 17517},
	{input: input{n: 3, p: []int{1, 2, 3}, q: []int{1, 2, 3}}, want: 0},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.p, tt.input.q)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	n := 8
	p := make([]int, n)
	for i, v := range rand.Perm(n) {
		p[i] = v + 1
	}
	rand.Seed(time.Now().UnixNano())
	q := make([]int, n)
	for i, v := range rand.Perm(n) {
		q[i] = v + 1
	}
	for i := 0; i < b.N; i++ {
		solution(n, p, q)
	}
}

func Example_main() {
	c, _ := ioutil.ReadFile("./input.txt")
	inr, inw, _ := os.Pipe()

	orgStdin := os.Stdin
	inw.Write(c)
	os.Stdin = inr
	defer func() {
		inw.Close()
		os.Stdin = orgStdin
	}()

	main()
	// Output: 17517
}
