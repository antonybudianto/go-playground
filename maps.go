package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	tmp := strings.Fields(s)
	
	mymap := make(map[string]int)
	for _, k := range tmp {
		mymap[k]++
	}
	return mymap
}

func main() {
	wc.Test(WordCount)
}
