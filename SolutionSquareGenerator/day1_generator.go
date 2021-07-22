package Smirnov

import "fmt"

func SolutionSquareGenerator(start int, n int) []int {
	var squares []int
	for i := start; i <= n; i++ {
		squares = append(squares, i*i)
	}
	return squares
}

func main() {
	fmt.Println(SolutionSquareGenerator(1, 10))
}
