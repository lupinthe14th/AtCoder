package main

import "testing"

var cases = []struct {
	id int
	input []int
	want  string
}{
	{id: 1, intput: []int{3, 3, 1, 1}, want: ""},
}

func TestMatrix(t *testing.T) {
	for _, tt:= range cases{
		t.Run(fmt.Sprint(tt.id), func(t *testing.T){
			got := ()
			if got != tt.want{
				t.Errorf("%s, want: %s", got , tt.want)
			}
		})
	}
}
