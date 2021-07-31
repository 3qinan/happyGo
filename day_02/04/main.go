package main

import "fmt"

//数组
func main() {

	var a1 [3]bool
	//1.初始化
	a1 = [3]bool{true, true, false}
	fmt.Println(a1)
	//2.初始化
	a10 := [...]int{1, 2, 3}
	fmt.Println(a10)
	//3.初始化
	a3 := [5]int{1, 2}
	fmt.Println(a3)
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4)
	//数组遍历
	citys := [...]string{"北京", "上海", "广州"}
	//1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//for range 遍历
	for i, v := range citys {
		fmt.Println(i, v)

	}
	//多维数组
	//var a11 [3][2]int
	a11 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)
	//多维数组遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	/*h := [5]int{1, 3, 5, 7, 8}
	for i := 0; i < len(h); i++ {
		sum := 0
		for _, v := range h {
			sum = sum + v

		}
		fmt.Println(sum)*/
	//}

}
