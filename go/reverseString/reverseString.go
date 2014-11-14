package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入待反转的字符串:")
	reader := bufio.NewReader(os.Stdin)
	str, _, _ := reader.ReadLine()

	orgText := string(str)
	s := []rune(orgText)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Printf(string(s))
}
