package main

//递归、阶乘
import "fmt"

/*func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}*/

//上台阶
func T(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return T(n-1) + T(n-2)
}

func main() {
	/*var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))*/

	ret := T(100)
	fmt.Println(ret)
}
