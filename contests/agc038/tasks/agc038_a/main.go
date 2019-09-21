package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func Matrix(h, w, a, b int) ([][]int, error) {
	log.Printf("h: %d", h)
	log.Printf("w: %d", w)
	log.Printf("a: %d", a)
	log.Printf("b: %d", b)
	// Slice of Slice
	dary := make([][]int, h)
	for i := range dary {
		dary[i] = make([]int, w)
	}

	// 初期値で条件を満たすか確認
	acount := make(map[int]int)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			acount[dary[i][j]]++
			if math.Min(float64(acount[0]), float64(acount[1])) == float64(a) {
				return dary, nil
			}
		}
	}

	// 初期値で条件を満たさない場合、bluteforceで逐一条件を満たすか確認する
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			dary[i][j] = 1
			log.Print(dary)
			acount[dary[i][j]]++
			if math.Min(float64(acount[0]), float64(acount[1])) == float64(a) {
				return dary, nil
			}
		}
	}
	return nil, errors.New("NG")
}

func main() {
	var h, w, a, b int
	fmt.Scan(&h, &w, &a, &b)
	matrix, err := Matrix(h, w, a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(matrix)
		for _, l := range matrix {
			sl := make([]string, 0, w)
			for _, s := range l {
				sl = append(sl, strconv.Itoa(s))
			}
			fmt.Println(strings.Join(sl, ""))
		}
	}
}
