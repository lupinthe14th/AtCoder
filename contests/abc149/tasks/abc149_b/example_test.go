package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func Example_main() {
	s := strings.NewReader("2 3 3")

	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	inw.Write(b)
	inw.Close()
	orgStdin := os.Stdin

	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output:
	// 0 2
}
