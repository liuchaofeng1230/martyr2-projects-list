package main

import "testing"

func TestCountWords(t *testing.T) {
	i := CountWords()

	if i != 11 {
		t.Log("统计元音字母数量错误")
		t.Fail()
	}
}
