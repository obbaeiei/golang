package golang

import "strconv"

func text(x int) string {
	if x == 3 {
		return "fizz"
	}
	return strconv.Itoa(x)
}
