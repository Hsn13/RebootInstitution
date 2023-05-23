package piscine

func ToLower(s string) string {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		c := str[i]
		if 'A' <= c && c <= 'Z' {
			c += 32
		}
		str[i] = c
	}
	return string(str)
}
