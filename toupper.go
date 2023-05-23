package piscine

func ToUpper(s string) string {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		c := str[i]
		if 'a' <= c && c <= 'z' {
			c -= 32
		}
		str[i] = c
	}
	return string(str)
}
