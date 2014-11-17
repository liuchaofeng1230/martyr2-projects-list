package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("请输入字符串:")
	reader := bufio.NewReader(os.Stdin)
	str, _, _ := reader.ReadLine()
	input := string(str)
	s := []rune(input)
	vowels := []string{"a", "e", "r", "o", "u"}
	store := make(map[string]int)

	// collect vowels
	for _, v := range s {
		c := string(v)
		if stringInSlice(c, vowels) {
			// if char is vowel
			if store[c] < 1 {
				store[c] = 0
			}
			store[c] += 1
		}
	}

	// count vowels
	for c, i := range store {
		fmt.Printf(c + ":" + strconv.Itoa(i) + "\n")
	}
}
