package rotationalcipher

func shiftFrom(key int, from rune) func(rune) rune {
	return func(l rune) rune {
		return rune((int(l-from)+key)%26) + from
	}
}

func shiftFn(shiftKey int) func(rune) rune {
	lowercaseShift := shiftFrom(shiftKey, 'a')
	uppercaseShift := shiftFrom(shiftKey, 'A')
	return func(l rune) rune {
		if 'a' <= l && l <= 'z' {
			return lowercaseShift(l)
		}
		if 'A' <= l && l <= 'Z' {
			return uppercaseShift(l)
		}
		return l
	}
}

func RotationalCipher(plain string, shiftKey int) string {
	shift := shiftFn(shiftKey)
	secret := []rune{}
	for _, v := range plain {
		secret = append(secret, shift(v))
	}
	return string(secret)
}
