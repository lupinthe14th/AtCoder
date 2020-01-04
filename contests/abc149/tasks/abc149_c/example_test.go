package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func Example_main() {
	s := strings.NewReader("20")

	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	inw.Write(b)
	inw.Close()
	orgStdin := os.Stdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: 23
}
