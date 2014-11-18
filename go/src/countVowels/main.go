package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func CountVowels(str string) map[string]int {
	vowels := []string{"a", "e", "i", "o", "u"}
	store := make(map[string]int)

	// collect vowels
	for _, v := range vowels {
		c := string(v)
		store[c] = strings.Count(str, c)
	}
	return store
}

func main() {
	input := utils.StdIn("请输入字符串:")
	store := CountVowels(input)

	// count vowels
	for c, i := range store {
		fmt.Printf(c + ":" + strconv.Itoa(i) + "\n")
	}
}
