package golang

import "strconv"

func text(x int) string {
	if x == 3 {
		return "fizz"
	}
	if x == 5 {
		return "buzz"
	}
	return strconv.Itoa(x)
}
