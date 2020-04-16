package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Perm struct {
	num int
	idx int
}

type Perms []Perm

func (p Perms) Len() int {
	return len(p)
}

func (p Perms) Less(i, j int) bool {
	return p[i].num < p[j].num
}

func (p Perms) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func swaps(n int, a, b []int) bool {
	ap := make(Perms, n)
	bp := make(Perms, n)
	for i := 0; i < n; i++ {
		ap[i] = Perm{idx: i, num: a[i]}
		bp[i] = Perm{idx: i, num: b[i]}
	}
	sort.Sort(ap)
	sort.Sort(bp)

	for i := 0; i < n; i++ {
		if ap[i].num > bp[i].num {
			return false
		}
	}

	for i := 0; i < n-1; i++ {
		if ap[i+1].num <= bp[i].num {
			return true
		}
	}

	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[ap[i].idx] = bp[i].idx
	}

	s := make(map[int]struct{})
	for i, m := 0, 0; i < n; i++ {
		s[m] = struct{}{}
		m = p[m]
	}

	if len(s) < n {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	n, err := strconv.Atoi(readLine(reader))
	checkError(err)
	a := make([]int, n)
	ss := strings.Split(readLine(reader), " ")
	for i := range ss {
		a[i], err = strconv.Atoi(ss[i])
		checkError(err)
	}

	b := make([]int, n)
	ss = strings.Split(readLine(reader), " ")
	for i := range ss {
		b[i], err = strconv.Atoi(ss[i])
		checkError(err)
	}

	if swaps(n, a, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
