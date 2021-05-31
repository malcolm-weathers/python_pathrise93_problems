package main

import (
	"fmt"
	"strings"
)

func combos(s string) []string {
	ret := []string{}
	if len(s) == 1 {
		ret = append(ret, strings.ToLower(string(s[0])))
		ret = append(ret, strings.ToUpper(string(s[0])))
	} else {
		recursive := combos(s[1:])
		for _, s2 := range recursive {
			ret = append(ret, strings.ToLower(string(s[0])) + s2)
			ret = append(ret, strings.ToUpper(string(s[0])) + s2)
		}
	}
	return ret
}

func main() {
	fmt.Println(combos("ab"))
}