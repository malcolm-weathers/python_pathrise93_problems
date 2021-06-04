package main

import "fmt"

func sum(x []int) int {
	s := 0
	for _, v := range(x) {
		s += v
	}
	return s
}

func maxsum(x []int) []int {
	if len(x) < 2 {
		return x
	}
	o1 := []int{x[0]}
	o2 := []int{x[1]}
	if len(x) > 2 {
		o1 = append(o1, maxsum(x[2:])...)
	}
	if len(x) > 3 {
		o2 = append(o2, maxsum(x[3:])...)
	}

	if sum(o1) > sum(o2) {
		return o1
	}
	return o2
}

func maxsum_dyn(x []int) []int {
	dyn := make([][]int, len(x))
	best := make([]int, len(x))
	
	dyn[len(dyn)-1] = []int{x[len(x)-1]}
	best[len(dyn)-1] = x[len(x)-1]
	dyn[len(dyn)-2] = []int{x[len(x)-2]}
	best[len(dyn)-2] = x[len(x)-2]
	dyn[len(dyn)-3] = []int{x[len(x)-1], x[len(x)-3]}
	best[len(dyn)-3] = x[len(x)-1] + x[len(x)-3]
	
	for i := 4; i <= len(x); i++ {
		f1 := i - 2
		f2 := i - 3
		o1 := len(x) - f1
		o2 := len(x) - f2
		oi := len(x) - i

		if best[o1] >= best[o2] {
			best[oi] = x[oi] + best[o1]
			dyn[oi] = append(dyn[o1], x[oi])
		} else {
			best[oi] = x[oi] + best[o2]
			dyn[oi] = append(dyn[o2], x[oi])
		}
	}
	
	best_index := 0
	for i := 1; i < len(best); i++ {
		if best[i] > best[best_index] {
			best_index = i
		}
	}
	
	for i, j := 0, len(dyn[best_index])-1; i < j; i, j = i+1, j-1 {
		dyn[best_index][i], dyn[best_index][j] = dyn[best_index][j], dyn[best_index][i]
	}
	
	return dyn[best_index]
}

func main() {
	fmt.Println(maxsum([]int{1, 2, 3}))
	fmt.Println(maxsum_dyn([]int{1, 2, 3}))
	fmt.Println(maxsum([]int{5, 1, 2, 6}))
	fmt.Println(maxsum_dyn([]int{5, 1, 2, 6}))
	fmt.Println(maxsum([]int{5, 1, 2, 6, 20, 2}))
	fmt.Println(maxsum_dyn([]int{5, 1, 2, 6, 20, 2}))
}
