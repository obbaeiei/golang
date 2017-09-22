package golang

func enc(xx, rot byte) byte {
	if xx == 32 {
		return xx
	}
	return ((xx + rot) % 26) + 97
}

func caesar(s string, rot byte) string {
	xx := []byte{}
	for i := range s {
		xx = append(xx, enc(s[i], rot))
	}

	return string(xx)
}
