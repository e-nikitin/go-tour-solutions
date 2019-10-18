package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, v := range strings.Fields(s) {
		m[v]++
	}

	return m
}

func main() {
	wc.Test(wordCount)
}
