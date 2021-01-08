package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	
	strField := strings.Fields(s)
	
	for _, c := range strField {
		m[string(c)] += 1
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}

