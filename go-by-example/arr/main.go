package main

import "fmt"

var arr [3]string

func main() {
	arr[0] = "go"
	arr[1] = "by"
	arr[2] = "example"

	for k, v := range arr {
		fmt.Println(k, "--->", v)
	}

	arr1 := [...]int32{1, 2}
	var arr2 [3]bool
	var arr3 [2]interface{}
	var arr4 [2]float64

	fmt.Println(arr1, arr2, arr3, arr4)

	fmt.Println(len(arr), cap(arr))
	fmt.Println(len(arr1), cap(arr1))
	fmt.Println(len(arr2), cap(arr2))

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println(c)

}
