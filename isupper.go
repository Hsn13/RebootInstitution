package piscine

func IsUpper(s string) bool {
	for _, s := range s {
		if s < 65 || s > 90 {
			return false
		}
	}
	return true
}
