package golang

func factorial(num int) int {
	var result = 1
	for ; 0 < num; num-- {
		result = result * num
	}
	return result
}
