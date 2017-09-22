package golang

func loop(top int) {
	i := 0
	for {
		i++
		if i > top {
			return
		}
	}
}
