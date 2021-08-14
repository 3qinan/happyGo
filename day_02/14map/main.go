package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1) //还没有初始化
	m1 = make(map[string]int, 10)
	m1["理想"] = 18
	m1["卢聪"] = 35
	fmt.Println(m1)
	fmt.Println(m1["理想"]) //取理想的值
	value, ok := m1["理想"]
	if !ok {
		fmt.Println("查无此key")

	} else {
		fmt.Println(value)
	}
	//map遍历k,v
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//只遍历k
	for k := range m1 {
		fmt.Println(k)
	}
	//只遍历v
	for _, v := range m1 {
		fmt.Println(v)
	}
	//删除
	delete(m1, "理想")
	delete(m1, "沙河") //删除不存在的，
}
