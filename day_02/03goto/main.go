package main

import "fmt"

func main() {
	var flag = true
	for i := 0; i < 10; i++ {
		for j := 'a'; j < 'z'; j++ {
			if j == 'f' {
				break
				flag = true

			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break
		}
	}

}
