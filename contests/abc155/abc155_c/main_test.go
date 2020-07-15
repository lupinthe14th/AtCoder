package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

type in struct {
	n     int
	texts []string
}

type Case struct {
	in   in
	want []string
}

var cases = []Case{
	{in: in{n: 7, texts: []string{"beat", "vet", "beet", "bed", "vet", "bet", "beet"}}, want: []string{"beet", "vet"}},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.texts)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("in: %+v, got: %s, want: %s", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	c, _ := ioutil.ReadFile("./input.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(c)
	inw.Close()
	os.Stdin = orgStdin
	defer func() { os.Stdin = orgStdin }()
	os.Stdin = inr

	main()

	// Output: buffalo
}

func Example_main2() {
	c, _ := ioutil.ReadFile("./input2.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(c)
	inw.Close()
	os.Stdin = orgStdin
	defer func() { os.Stdin = orgStdin }()
	os.Stdin = inr

	main()

	// Output: kick
}

func Example_main3() {
	c, _ := ioutil.ReadFile("./input3.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(c)
	inw.Close()
	os.Stdin = orgStdin
	defer func() { os.Stdin = orgStdin }()
	os.Stdin = inr

	main()

	// Output:
	// kun
	// nichia
	// tapu
	// ushi

}
