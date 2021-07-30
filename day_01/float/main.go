package main

import "fmt"

func main() {
	f1 := 1.23456
	//默认float64
	fmt.Printf("%T\n", f1)
	f2 := float32(2.3445)
	fmt.Printf("%T\n", f2)
	f2 = float32(f1)

	//布尔值
	b1 := true
	var b2 bool
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T      value:%v\n", b2, b2)

}
