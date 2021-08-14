package main

import "fmt"

func main() {
	//append 补充

	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15}
	s1 := a1[:]
	s1 = append(s1[0:1], s1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1)
}
