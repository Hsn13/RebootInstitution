package piscine

func IsNumeric(s string) bool {
	for _, s := range s {
		if s < 48 || s > 57 {
			return false
		}
	}
	return true
}
