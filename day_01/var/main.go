package main

import "fmt"

var name string
var age int

func foo() (int, string) {
	return 10, "cxk"
}

func main() {
	//声明变量
	name = "蔡徐坤"
	age = 24
	var s1 = "大傻逼"
	s2 := "是真的"
	fmt.Printf("name:%s", name)
	fmt.Println()
	fmt.Println(age)
	fmt.Println(s1)
	fmt.Println(s2)

	x, _ := foo()
	_, y := foo()
	fmt.Println(x)
	fmt.Println(y)

}
