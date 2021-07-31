package main

import "fmt"

func main() {
	//切片定义
	var s1 []int //定义一个存放int的切片
	var s2 []string
	fmt.Println(s1, s2)
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"上", "下", "左"}
	fmt.Println(s1, s2)
	//长度和容量
	fmt.Printf("len(s1):%d\n cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d\n cap(s1):%d\n", len(s2), cap(s2))
	//2.将数组切割到切片
	a1 := [...]int{1, 2, 34, 56, 6}
	s3 := a1[0:4] //左包含，右不包含【1，2，34，56】
	s4 := a1[:]
	s5 := a1[:4]
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

}
