package main

import (
	"fmt"
)

/*func 函数名(参数)(返回值){
	函数体
}*/
func intSum(x int, y int) int {
	return x + y
}

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)

}

//没有参数 没有返回值
func f2() {
	fmt.Println("f2")
}

//没有参数,返回值没命名
func f3() int {
	ret := 3
	return ret
}
func f4(x int, y int) (ret int) {
	ret = x + y
	return //使用命名返回值可以省略return后面内容
}

//多个返回值
func f5() (int, string) {
	return 1, "lp"
}

func f6(x, y int) int { //多参数类型相同，可以省略前面参数类型
	return x + y
}

//可变长参数
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)

}
func main() {

	r := intSum(1, 3)
	fmt.Println(r)
	f7("李鹏")
	f7("lpp", 1, 2, 3)
}
