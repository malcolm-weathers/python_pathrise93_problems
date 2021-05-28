package main

import "fmt"

func main() {
	arr := []int{5, 9, 0, -11, -1, 5, 17, 8, 117, 5}
	dup := []int{}
	for _, v := range arr {
		is_in_dup := false;
		for _, x := range dup {
			if v == x {
				is_in_dup = true;
			}
		}
		if !is_in_dup {
			dup = append(dup, v)
		}
	}
	
	for _, v := range dup {
		fmt.Printf("%d ", v);
	}
	fmt.Println("")
}
