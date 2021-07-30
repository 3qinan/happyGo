package main

import "fmt"

func main() {
	var i1 = 101
	i2 := 077
	i3 := 0xf
	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3)
	//查看变量类型 %T
	fmt.Printf("%T\n", i3)

	//声明一int8 类型的变量
	i4 := int8(10)
	fmt.Printf("%T", i4)

}
