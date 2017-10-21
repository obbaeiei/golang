package golang

func factorial(num int) int {
	var result = 1
	for 0 < num {
		result = result * num
		num--
	}
	return result
}
