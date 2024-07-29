package main

import (
	"fmt"
	"testing"
)

// go test -v .
func TestIntMin(t *testing.T) {
	rlt := IntMin(23, -20)
	if rlt != -200 {
		t.Errorf("expect -20, but is  %d", rlt)
	}
}

// go test -bench=. -run=^IntMin$
//
//	go test -bench=. -run=^$
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(20, 30)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	casees := []struct {
		a, b int
		want int
	}{
		{1, 4, 1},
		{0, 1, 0},
		{-1, 9, -1},
		{-2, 8, -20},
		{-199, -10, -199},
		{99, 100, 99},
	}

	for _, v := range casees {
		name := fmt.Sprintf("案例是: %d,%d", v.a, v.b)

		t.Run(name, func(t *testing.T) {
			rlt := IntMin(v.a, v.b)
			if rlt != v.want {
				//t.Errorf("wants %d but get %d", v.want, rlt)
				t.Log("还能执行")
				t.Fatalf("wants %d but get %d", v.want, rlt)
				//fmt.Println("test")
				t.Log("还能执行吗？？")

			}
		})
	}
}
