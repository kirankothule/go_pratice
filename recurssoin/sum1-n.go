package main

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}
