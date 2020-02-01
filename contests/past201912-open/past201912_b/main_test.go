package main

import (
	"io/ioutil"
	"os"
)

func Example_main() {
	c, _ := ioutil.ReadFile("./input.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(c)
	inw.Close()
	os.Stdin = orgStdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output:
	// up 1
	// down 7
	// up 97
	// stay
	// down 10
	// down 10
	// down 70
	// up 20
	// down 20
}
