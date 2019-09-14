package abc087b

import "fmt"

func coins(a, b, c, x int) (p int) {
	p = 0
	switch {
	case x >= 500:
		for i := 0; a >= i; i++ {
			for j := 0; b >= j; j++ {
				for k := 0; c >= k; k++ {
					if 500*i+100*j+50*k == x {
						p++
					}
				}
			}

		}
	case x >= 100:
		for j := 0; b >= j; j++ {
			for k := 0; c >= k; k++ {
				if 100*j+50*k == x {
					p++
				}
			}
		}
	case x >= 50:
		for k := 0; c >= k; k++ {
			if 50*k == x {
				p++
			}
		}

	}
	if x < 50 {
		panic("fail")
	}
	return
}

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	fmt.Println(coins(a, b, c, x))
}
