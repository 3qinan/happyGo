package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//判断汉字数量
	s1 := "hello李宁波"
	var count = 0
	for _, c := range s1 { //依次拿到字符
		if unicode.Is(unicode.Han, c) { //判断是汉字
			count++

		}
	}
	fmt.Println(count)
	//判断单词出现次数
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")   //按照空格切割s2
	m1 := make(map[string]int, 10) //遍历切片存到map
	for _, c := range s3 {
		//如果原来map中不存在那么出现次数=1
		//如果存在，次数+1
		if _, ok := m1[c]; !ok {
			m1[c] = 1
		} else {
			m1[c]++
		}
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}
