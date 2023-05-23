package piscine

func BasicJoin(elems []string) string {
	res := ""
	for _, val := range elems {
		res = res + val
	}
	return res
}
