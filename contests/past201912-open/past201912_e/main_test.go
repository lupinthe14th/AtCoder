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
	// NYYNYY
	// NNYNNN
	// NNNYNN
	// NNNNNN
	// NNNNNY
	// YNNNYN
}
