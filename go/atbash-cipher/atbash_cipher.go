package atbash

func Atbash(plain string) string {
	secret := []rune{}
	c := 0
	for _, v := range plain {
		if '0' <= v && v <= '9' {
			if c == 5 {
				secret = append(secret, ' ')
				c = 0
			}
			secret = append(secret, v)
			c++
		} else if 'a' <= v && v <= 'z' {
			if c == 5 {
				secret = append(secret, ' ')
				c = 0
			}
			secret = append(secret, rune('a'+'a'+25-int(v)))
			c++
		} else if 'A' <= v && v <= 'Z' {
			if c == 5 {
				secret = append(secret, ' ')
				c = 0
			}
			secret = append(secret, rune('A'+'a'+25-int(v)))
			c++
		}

	}
	return string(secret)
}
