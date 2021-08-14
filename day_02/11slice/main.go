package main

import "fmt"

func main() {
	var a = make([]int, 5, 10) //创建的时候已经是【00000】int类型默认为值"0"
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

}
