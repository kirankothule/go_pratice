package main

import (
	"fmt"
)

func main() {
	const (
		a = iota // initilize with 0
		b
		c
		d
		e
	)
	fmt.Println(a, b, c, d, e)

	const (
		p = iota // The black identifier used to give the value to const alternate number
		_
		q
		_
		r
	)
	fmt.Println(p, q, r)

}

// o/p : 0 1 2 3 4
