package main

import (
	"fmt"
	ioutil "io/ioutil"
	"strconv"
	"strings"
)

func CountWords() int {
	f, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("读文件出错")
	}

	str := string(f)
	return len(strings.Fields(str))
}

func main() {
	length := CountWords()
	fmt.Println(strconv.Itoa(length))
}
