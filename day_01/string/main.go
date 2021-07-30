package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "\"d:\\go\\src\\day_01\""
	fmt.Println(path)

	s1 := `
     我是小学生
       大碗宽面  
   不尽长江滚滚来
  `
	fmt.Println(s1)
	//字符串长度
	fmt.Println(len(s1))
	//字符串拼接
	na := "lnb"
	wo := "666"
	ss := na + wo
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", na, wo)
	fmt.Sprintf("%s%s", na, wo)
	fmt.Println(ss1)
	//分割字符串
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(ss, "lnb"))
	//前缀
	fmt.Println(strings.HasPrefix(ss, "ln"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "6"))

	s7 := "agscjgsa"

	fmt.Println(strings.Index(s7, "j"))
	fmt.Println(strings.LastIndex(s7, "a"))
	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
