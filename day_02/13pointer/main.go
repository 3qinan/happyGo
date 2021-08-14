package main

import "fmt"

func main() {
	//指针
	//1.&:取地址
	//2.*:根据地址取值
	n := 18
	p := &n
	fmt.Println(p)        //指针内存地址
	fmt.Printf("%T\n", p) //指针类型
	//2.
	m := *p
	fmt.Println(m)
	//new函数申请内存地址,基本型 int ，string...等，返回的是对应类型的指针
	var a = new(int)
	*a = 100
	fmt.Println(*a)
	//make用于 slice、map、chan的内存地址申请，返回的是类型本身
}
