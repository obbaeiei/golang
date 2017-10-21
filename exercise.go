package golang

func concat(data, add []string) []string {
	return append(data, add...)
}

func delFirst(data []string) []string {
	return data[1:]
}

func delLast(data []string) []string {
	return data[0:(len(data) - 1)]
}

func delSecond(data []string) []string {
	return concat(data[0:1], data[2:])
}

func odd(data []int) []int {
	// filter
	for i := 0; i < len(data); i++ {

	}
	return data
}
