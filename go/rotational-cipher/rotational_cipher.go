package rotationalcipher



func shiftFn(shiftKey int) func(rune) rune {
	return func(l rune) rune {
		if('a' <= l && l <= 'z') {
			return (rune((int(l-'a') + shiftKey)%26) + 'a')
		}
		if('A' <= l && l <= 'Z') {
			return (rune((int(l-'A') + shiftKey)%26) + 'A')
		}
		return l
	}
}

func RotationalCipher(plain string, shiftKey int) string {
	shift := shiftFn(shiftKey)
	secret := []rune{}
	for _, v:= range plain {
		secret = append(secret, shift(v))
	}
	return string(secret)
}
