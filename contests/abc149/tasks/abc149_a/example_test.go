package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func Example_main() {
	s := strings.NewReader("oder atc")
	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output:
	// atcoder
}
