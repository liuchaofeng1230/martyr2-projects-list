package utils

import (
	"bufio"
	"fmt"
	"os"
)

func StdIn(tip string) (str string) {
	fmt.Println(tip)
	reader := bufio.NewReader(os.Stdin)
	buf, _, _ := reader.ReadLine()

	str = string(buf)
	return
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
