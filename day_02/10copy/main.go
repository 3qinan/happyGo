package main

import "fmt"

func main() {
	//copy
	a1 := []int{1, 3, 5, 7}
	//a2 := a1 //赋值
	/*	var a3 = make([]int, 3, 3)
		copy(a3, a1)
		fmt.Println(a1, a2, a3)
		a1[0] = 100
		fmt.Println(a1, a2, a3)*/
	//将a1中索引为1的元素"3"删掉
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)
	fmt.Println(cap(a1))
	a1[0] = 100
	fmt.Println(a1)
}
