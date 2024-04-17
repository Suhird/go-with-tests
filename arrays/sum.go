package main

func Sum(n []int) int {
	sum := 0
	for _, num := range n {
		sum += num
	}
	return sum
}
