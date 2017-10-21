package golang

func fibonacci(x int) []int {
	results := []int{}
	for i := 0; i < x; i++ {
		if i == 0 {
			results = append(results, i)
		} else if i == 1 {
			results = append(results, i+results[i-1])
		} else {
			results = append(results, results[i-2]+results[i-1])
		}
	}
	return results
}
