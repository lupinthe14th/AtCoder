package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func Example_main() {
	s := strings.NewReader("oder atc")
	c, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write([]byte(c))
	inw.Close()
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output:
	// atcoder
}
