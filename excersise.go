package golang

func factorial(num int) int {
	var result = 1
	for i := 0; i < num; i++ {
		result = result * (num - i)
	}
	return result
}
