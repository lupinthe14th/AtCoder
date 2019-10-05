package agc039a

import (
	"fmt"
)

func connectionAndDisconnection(s string, k int) int {

	r := []rune(s)
	counter := 0
	j := 0
	for i := 1; i < len(r); i++ {
		if r[j] == r[i] {
			counter++
			r[i] = r[i] + 1
		}
		j++
	}
	return counter * k
}

func main() {
	var s string
	var k int
	fmt.Scan(&s)
	fmt.Scan(&k)
	fmt.Println(connectionAndDisconnection(s, k))
}
