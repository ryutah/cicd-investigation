package main

func factorial(x int) int {
	if x <= 0 {
		return 1
	}
	return factorial(x-1) * x
}
