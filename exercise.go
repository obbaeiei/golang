package golang

func enc(x rune, rot byte) rune {
	return ((int32(x) + int32(rot)) % 26) + (int32([]rune("a")[0]))
}

//
func caesar(s string, rot byte) string {
	var result string
	pieceOfString := []rune(s)
	for _, v := range pieceOfString {
		if v != int32((int32([]rune(" ")[0]))) {
			result += string(enc(v, rot))
		} else {
			result += " "
		}
	}
	return result
}
