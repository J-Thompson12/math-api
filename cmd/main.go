package main

import (
	"fmt"
	"strings"
)

func main() {
	x := compare("0.1.1.1", "0.2.0.1.2")
	fmt.Println(x)
}

func compare(v1, v2 string) int {
	if v1 == v2 {
		return 0
	}

	v1Arr := strings.Split(v1, ".")
	v2Arr := strings.Split(v2, ".")

	matchArrSize(&v1Arr, &v2Arr)

	for i, num := range v1Arr {
		if num < v2Arr[i] {
			return -1
		}
		if num > v2Arr[i] {
			return 1
		}

	}
	return 0
}

// make the versions the same length to avoid nil element
func matchArrSize(arr1, arr2 *[]string) {
	if len(*arr1) == len(*arr2) {
		return
	}
	if len(*arr1) > len(*arr2) {
		for i := range *arr1 {
			if i+1 > len(*arr2) {
				*arr2 = append(*arr2, "0")
			}
		}
	} else {
		for i := range *arr2 {
			if i+1 > len(*arr1) {
				*arr1 = append(*arr1, "0")
			}
		}
	}
}
