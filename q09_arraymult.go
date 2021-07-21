package main

import "fmt"

func simple(arr []int) []int {
	ret := []int{}
	num := 1
	for _, n := range(arr) {
		num *= n
	}
	for _, n := range(arr) {
		ret = append(ret, num/n)
	}
	return ret
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ret := simple(arr)
	for _, n := range(ret) {
		fmt.Print(n, " ");
	}
}
