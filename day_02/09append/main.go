package main

import "fmt"

func main() {
	//append()追加元素
	s1 := []string{"北京", "上海", "广州"}
	s1 = append(s1, "徐州")
	fmt.Println(s1)

}
