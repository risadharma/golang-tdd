package main

func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	return sum
}

func main() {}
