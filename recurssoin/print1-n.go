package main

import "fmt"

func main() {
	print(10)
	fmt.Println(sum(10))
}

func print(n int) {
	if n == 0 {
		return
	}
	print(n - 1)
	fmt.Println(n)
}
