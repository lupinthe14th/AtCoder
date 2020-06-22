package main

import (
	"fmt"
	"sort"
)

type restaurant struct {
	id, rate int
	city     string
}

func solution(n int, restaurants []restaurant) []int {
	sort.SliceStable(restaurants, func(i, j int) bool {
		if restaurants[i].city == restaurants[j].city {
			return restaurants[i].rate > restaurants[j].rate
		}
		return restaurants[i].city < restaurants[j].city
	})

	out := make([]int, n)
	for i := range restaurants {
		out[i] = restaurants[i].id
	}
	return out
}

func main() {
	var n int
	fmt.Scan(&n)
	restaurants := make([]restaurant, n)
	for i := range restaurants {
		restaurants[i].id = i + 1
		fmt.Scan(&restaurants[i].city, &restaurants[i].rate)
	}
	for _, t := range solution(n, restaurants) {
		fmt.Println(t)
	}
}
