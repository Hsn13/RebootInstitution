package piscine

func SplitWhiteSpaces(s string) []string {
	runes := []rune(s)
	counter := 0
	for _, val := range runes {
		if val == ' ' || val == '\t' || val == '\n' {
			counter++
		}
	}
	i := 0
	res_arr := make([]string, counter)
	tmp := ""
	for _, val := range runes {
		if val == ' ' || val == '\t' || val == '\n' {
			res_arr[i] = tmp
			tmp = ""
			i++
		} else {
			tmp += string(val)
		}
	}
	return res_arr
}
