package main

import (
	"log"
	"slices"
)

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return fib(n-2) + fib(n-1)
	}
}
func fact(n int) int {
	//或者 n==0
	if n <= 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

// 递归形式
func binarySearch(in []int, target int) int {
	if len(in) == 0 {
		return -1
	}
	start := 0
	end := len(in) - 1
	mid := (start + end) / 2
	if target == in[mid] {
		return mid
	} else if target > in[mid] {
		start = mid + 1
		return binarySearch(in[mid+1:], target)
	} else {
		end = mid - 1
		return binarySearch(in[0:mid], target)
	}
}

// 非递归形式
func binarySearchNoCurse(in []int, target int) int {
	if len(in) <= 0 {
		return -1
	}
	start := 0
	end := len(in) - 1
	for start <= end {
		mid := (end + start) / 2
		if target > in[mid] {
			start = mid + 1
		} else if target < in[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
func main() {
	var input = []int{4, 5, 6, 11, 9, 4, 5, 20, 1}
	log.Println(input)

	//sort2.Strings()
	slices.Sort(input)
	log.Println(input)

	//sort2.Strings("")
	//sort2.Reverse()

	//log.Println(binarySearchNoCurse(input, 11))
	log.Println(binarySearch(input, 20))

	var f3 func(int) int
	f3 = func(n int) int {
		if n <= 1 {
			return 1
		} else {
			return n * f3(n-1)
		}
	}
	log.Println(fib(7))

	log.Println(f3(7))

}
